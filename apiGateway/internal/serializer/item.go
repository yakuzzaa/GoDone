package serializer

import "time"

type ItemInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type Item struct {
	Id       uint64     `json:"id"`
	Info     *ItemInfo  `json:"info"`
	CreateAt *time.Time `json:"create_at"`
	UpdateAt *time.Time `json:"update_at"`
}

type UpdateItemInfo struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}
type CreateItemRequest struct {
	Info ItemInfo `json:"info"`
}

type CreateItemResponse struct {
	Id uint64 `json:"id"`
}

type ListItemResponse struct {
	Items []Item `json:"items"`
}

type GetItemResponse struct {
	Item Item `json:"item"`
}

type UpdateItemRequest struct {
	Info UpdateItemInfo `json:"info"`
}
type UpdateItemResponse struct {
	Status string `json:"status"`
}

type DeleteItemResponse struct {
	Status string `json:"status"`
}
