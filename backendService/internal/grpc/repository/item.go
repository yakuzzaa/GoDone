package repository

import (
	"errors"

	"github.com/yakuzzaa/GoDone/backendService/grpc/pkg/item_v1"
	"github.com/yakuzzaa/GoDone/backendService/internal/models"
	"gorm.io/gorm"
)

type ItemRepositoryInterface interface {
	List(listId uint64) ([]models.Item, error)
	GetById(itemId uint64) (*models.Item, error)
	Create(listId uint64, itemInfo *item_v1.ItemInfo) (uint64, error)
	Update(itemId uint64, itemInfo *item_v1.UpdateItemInfo) error
	Delete(itemId uint64) error
}

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepositoryInterface {
	return &ItemRepository{db: db}
}

func (i *ItemRepository) List(listId uint64) ([]models.Item, error) {
	var items []models.Item
	if err := i.db.Where("to_do_list_id = ?", listId).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (i *ItemRepository) GetById(itemId uint64) (*models.Item, error) {
	var item models.Item
	if err := i.db.Preload("Items").Where("id = ?", itemId).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (i *ItemRepository) Create(listId uint64, itemInfo *item_v1.ItemInfo) (uint64, error) {
	item := models.Item{
		Title:       itemInfo.Title,
		Description: itemInfo.Description,
		ToDoListID:  uint(listId),
		Done:        false,
	}

	if err := i.db.Create(&item).Error; err != nil {
		return 0, err
	}

	return uint64(item.ID), nil
}

func (i *ItemRepository) Update(itemId uint64, itemInfo *item_v1.UpdateItemInfo) error {
	item, err := i.IsItemExist(itemId)
	if err != nil {
		return err
	}
	if itemInfo.Title != nil {
		item.Title = itemInfo.Title.Value
	}
	if itemInfo.Description != nil {
		item.Description = itemInfo.Description.Value
	}

	if err := i.db.Save(&item).Error; err != nil {
		return err
	}

	return nil
}

func (i *ItemRepository) Delete(itemId uint64) error {
	if err := i.db.Delete(&models.ToDoList{}, itemId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("item not found")
		}
		return err
	}
	return nil
}

func (i *ItemRepository) IsItemExist(itemId uint64) (*models.Item, error) {
	var item *models.Item
	if err := i.db.Where("id = ?", itemId).First(&item).Error; err != nil {
		return &models.Item{}, err
	}
	return item, nil
}
