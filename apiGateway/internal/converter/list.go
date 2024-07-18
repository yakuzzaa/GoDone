package converter

import (
	"time"

	"github.com/jinzhu/copier"
	"github.com/yakuzzaa/GoDone/apiGateway/internal/serializer"
	list "github.com/yakuzzaa/GoDone/backendService/grpc/pkg/list_v1"
)

func MarshalGetResponse(response *list.ListResponse) (*serializer.ListsResponse, error) {
	var lists []serializer.List
	for _, l := range response.Lists {
		var info serializer.ListInfo
		var createdAt time.Time
		var updatedAt time.Time
		if err := copier.Copy(&info, l.Info); err != nil {
			return nil, err
		}

		if err := copier.Copy(&createdAt, l.CreatedAt); err != nil {
			return nil, err
		}

		if err := copier.Copy(&updatedAt, l.UpdatedAt); err != nil {
			return nil, err
		}

		lists = append(lists, serializer.List{
			Id:        l.Id,
			Info:      &info,
			CreatedAt: &createdAt,
			UpdatedAt: &updatedAt,
		})
	}
	return &serializer.ListsResponse{Lists: lists}, nil
}

func MarshalDetailResponse(response *list.DetailResponse) (*serializer.ListByIdResponse, error) {
	var apiList serializer.List
	var apiItems []serializer.Item
	var createdAt time.Time
	var updatedAt time.Time

	if err := copier.Copy(&apiList, response.List.List); err != nil {
		return nil, err
	}

	if err := copier.Copy(&createdAt, response.List.List.CreatedAt); err != nil {
		return nil, err
	}
	apiList.CreatedAt = &createdAt

	if err := copier.Copy(&updatedAt, response.List.List.UpdatedAt); err != nil {
		return nil, err
	}
	apiList.UpdatedAt = &updatedAt

	if err := copier.Copy(&apiList.Info, response.List.List.Info); err != nil {
		return nil, err
	}

	for _, item := range response.List.Items {
		var apiItem serializer.Item
		if err := copier.Copy(&apiItem, item); err != nil {
			return nil, err
		}
		apiItems = append(apiItems, apiItem)
	}

	return &serializer.ListByIdResponse{
		List: serializer.ListWithItems{
			List:  apiList,
			Items: apiItems,
		},
	}, nil
}
