package service

import (
	"context"

	"github.com/yakuzzaa/GoDone/backendService/grpc/pkg/list_v1"
	"github.com/yakuzzaa/GoDone/backendService/internal/grpc/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ListServiceInterface interface {
	CreateList(listInfo *list_v1.ListInfo, userId uint64) (uint64, error)
	GetListById(id uint64) (*list_v1.List, []*list_v1.Item, error)
	GetList(userId uint64) (*[]*list_v1.List, error)
	Update(id uint64, updateInfo *list_v1.UpdateList) error
	Delete(ctx context.Context, id uint64) error
}
type ListService struct {
	repo repository.ListRepositoryInterface
}

func NewListService(repo repository.ListRepositoryInterface) ListServiceInterface {
	return &ListService{repo: repo}
}

func (l *ListService) CreateList(listInfo *list_v1.ListInfo, userId uint64) (uint64, error) {
	listId, err := l.repo.Create(listInfo, userId)
	if err != nil {
		return 0, err
	}
	return listId, nil
}

func (l *ListService) GetListById(id uint64) (*list_v1.List, []*list_v1.Item, error) {
	detail, err := l.repo.GetById(id)
	if err != nil {
		return nil, nil, err
	}
	protoList := &list_v1.List{
		Id: id,
		Info: &list_v1.ListInfo{
			Title:       detail.Title,
			Description: detail.Description,
		},
		CreatedAt: timestamppb.New(detail.CreatedAt),
		UpdatedAt: timestamppb.New(detail.UpdatedAt),
	}

	protoItems := make([]*list_v1.Item, 0, len(detail.Items))
	for _, item := range detail.Items {
		element := &list_v1.Item{
			Id: uint64(item.ID),
			Info: &list_v1.ItemInfo{
				Title:       item.Title,
				Description: item.Description,
				Done:        item.Done,
			},
			CreatedAt: timestamppb.New(item.CreatedAt),
			UpdatedAt: timestamppb.New(item.UpdatedAt),
		}
		protoItems = append(protoItems, element)
	}

	return protoList, protoItems, nil
}

func (l *ListService) GetList(userId uint64) (*[]*list_v1.List, error) {
	lists, err := l.repo.List(userId)
	if err != nil {
		return nil, err
	}

	protoLists := make([]*list_v1.List, 0, len(lists))

	for _, list := range lists {
		element := &list_v1.List{
			Id: uint64(list.ID),
			Info: &list_v1.ListInfo{
				Title:       list.Title,
				Description: list.Description,
			},
			CreatedAt: timestamppb.New(list.CreatedAt),
			UpdatedAt: timestamppb.New(list.UpdatedAt),
		}
		protoLists = append(protoLists, element)
	}

	return &protoLists, nil
}

func (l *ListService) Update(id uint64, updateInfo *list_v1.UpdateList) error {
	err := l.repo.Update(id, updateInfo)
	if err != nil {
		return err
	}
	return nil
}

func (l *ListService) Delete(ctx context.Context, id uint64) error {
	err := l.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
