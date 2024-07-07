package service

import (
	"github.com/yakuzzaa/GoDone/backendService/grpc/pkg/item_v1"
	"github.com/yakuzzaa/GoDone/backendService/internal/grpc/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ItemServiceInterface interface {
	CreateItem(listId uint64, itemInfo *item_v1.ItemInfo) (uint64, error)
	GetList(listId uint64) (*[]*item_v1.Item, error)
	GetById(itemId uint64) (*item_v1.Item, error)
	Update(itemId uint64, updateInfo *item_v1.UpdateItemInfo) error
	Delete(itemId uint64) error
}

type ItemService struct {
	repo repository.ItemRepositoryInterface
}

func NewItemRepository(repo repository.ItemRepositoryInterface) ItemServiceInterface {
	return &ItemService{repo: repo}
}

func (i *ItemService) CreateItem(listId uint64, itemInfo *item_v1.ItemInfo) (uint64, error) {
	itemId, err := i.repo.Create(listId, itemInfo)
	if err != nil {
		return 0, err
	}
	return itemId, nil
}

func (i *ItemService) GetList(listId uint64) (*[]*item_v1.Item, error) {
	items, err := i.repo.List(listId)
	if err != nil {
		return nil, err
	}

	protoItems := make([]*item_v1.Item, 0, len(items))

	for _, item := range items {
		element := &item_v1.Item{
			Id: uint64(item.ID),
			Info: &item_v1.ItemInfo{
				Title:       item.Title,
				Description: item.Description,
				Done:        item.Done,
			},
			CreatedAt: timestamppb.New(item.CreatedAt),
			UpdatedAt: timestamppb.New(item.UpdatedAt),
		}
		protoItems = append(protoItems, element)
	}

	return &protoItems, nil
}

func (i *ItemService) GetById(itemId uint64) (*item_v1.Item, error) {
	item, err := i.repo.GetById(itemId)
	if err != nil {
		return nil, err
	}
	protoItem := &item_v1.Item{
		Id: itemId,
		Info: &item_v1.ItemInfo{
			Title:       item.Title,
			Description: item.Description,
			Done:        item.Done,
		},
		CreatedAt: timestamppb.New(item.CreatedAt),
		UpdatedAt: timestamppb.New(item.UpdatedAt),
	}

	return protoItem, nil
}

func (i *ItemService) Update(itemId uint64, updateInfo *item_v1.UpdateItemInfo) error {
	err := i.repo.Update(itemId, updateInfo)
	if err != nil {
		return err
	}
	return nil
}

func (i *ItemService) Delete(itemId uint64) error {
	err := i.repo.Delete(itemId)
	if err != nil {
		return err
	}
	return nil
}
