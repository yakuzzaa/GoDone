package converter

import (
	"time"

	"github.com/jinzhu/copier"
	"github.com/yakuzzaa/GoDone/apiGateway/internal/serializer"
	item "github.com/yakuzzaa/GoDone/backendService/grpc/pkg/item_v1"
)

func MarshalGetItemResponse(response *item.ListResponse) (*serializer.ListItemResponse, error) {
	var items []serializer.Item
	for _, i := range response.Items {
		var info serializer.ItemInfo
		var createdAt time.Time
		var updatedAt time.Time
		if err := copier.Copy(&info, i.Info); err != nil {
			return nil, err
		}

		if err := copier.Copy(&createdAt, i.CreatedAt); err != nil {
			return nil, err
		}

		if err := copier.Copy(&updatedAt, i.UpdatedAt); err != nil {
			return nil, err
		}

		items = append(items, serializer.Item{
			Id:       i.Id,
			Info:     &info,
			CreateAt: &createdAt,
			UpdateAt: &updatedAt,
		})
	}
	return &serializer.ListItemResponse{Items: items}, nil
}

func MarshalDetailItemResponse(response *item.GetResponse) (*serializer.GetItemResponse, error) {
	var Item serializer.Item

	if err := copier.Copy(&Item, response.Item); err != nil {
		return nil, err
	}

	return &serializer.GetItemResponse{
		Item: Item}, nil
}
