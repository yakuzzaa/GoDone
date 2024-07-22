package repository

import (
	"context"
	"errors"

	"github.com/yakuzzaa/GoDone/backendService/grpc/pkg/list_v1"
	"github.com/yakuzzaa/GoDone/backendService/internal/models"
	"gorm.io/gorm"
)

type ListRepositoryInterface interface {
	Create(listInfo *list_v1.ListInfo, userId uint64) (uint64, error)
	List(userId uint64) ([]models.ToDoList, error)
	GetById(id uint64, userId uint64) (*models.ToDoList, error)
	Update(id uint64, userId uint64, listInfo *list_v1.UpdateList) error
	Delete(ctx context.Context, id uint64, userId uint64) error
}
type ListRepository struct {
	db *gorm.DB
}

func NewListRepository(db *gorm.DB) *ListRepository {
	return &ListRepository{db: db}
}

func (l *ListRepository) Create(listInfo *list_v1.ListInfo, userId uint64) (uint64, error) {
	list := models.ToDoList{
		UserID:      uint(userId),
		Title:       listInfo.Title,
		Description: listInfo.Description,
	}

	if err := l.db.Create(&list).Error; err != nil {
		return 0, err
	}

	return uint64(list.ID), nil
}

func (l *ListRepository) List(userId uint64) ([]models.ToDoList, error) {
	var lists []models.ToDoList
	if err := l.db.Where("user_id = ?", userId).Find(&lists).Error; err != nil {
		return nil, err
	}
	return lists, nil
}

func (l *ListRepository) GetById(id uint64, userId uint64) (*models.ToDoList, error) {
	var list models.ToDoList
	if err := l.db.Preload("Items").Where("id = ? AND user_id = ?", id, userId).First(&list).Error; err != nil {
		return nil, err
	}
	return &list, nil
}

func (l *ListRepository) Update(id uint64, userId uint64, listInfo *list_v1.UpdateList) error {
	list, err := l.IsListExist(id, userId)
	if err != nil {
		return err
	}
	if listInfo.Title != nil {
		list.Title = listInfo.Title.Value
	}
	if listInfo.Description != nil {
		list.Description = listInfo.Description.Value
	}

	if err := l.db.Save(&list).Error; err != nil {
		return err
	}

	return nil
}

func (l *ListRepository) Delete(ctx context.Context, id uint64, userId uint64) error {
	tx := l.db.WithContext(ctx).Begin()

	if err := tx.WithContext(ctx).Where("to_do_list_id = ? AND user_id = ?", id, userId).Delete(&[]models.Item{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.WithContext(ctx).Delete(&models.ToDoList{}, id).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("list not found")
		}
		return err
	}

	return tx.Commit().Error
}

func (l *ListRepository) IsListExist(id uint64, userId uint64) (*models.ToDoList, error) {
	var list *models.ToDoList
	if err := l.db.Where("id = ? AND user_id = ?", id, userId).First(&list).Error; err != nil {
		return &models.ToDoList{}, err
	}
	return list, nil
}
