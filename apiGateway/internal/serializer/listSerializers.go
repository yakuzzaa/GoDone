package serializer

import (
	"time"
)

type ListWithItems struct {
	List  List   `json:"lists"`
	Items []Item `json:"items"`
}

type ListInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateListInfo struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type List struct {
	Id        uint64     `json:"id"`
	Info      *ListInfo  `json:"info"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ListsRequest struct {
}
type ListsResponse struct {
	Lists []List `json:"lists"`
}

type ListByIdRequest struct {
	ListId uint64 `json:"id"`
	UserId uint64 `json:"_"`
}

type ListByIdResponse struct {
	List ListWithItems `json:"list"`
}

type CreateListRequest struct {
	ListInfo ListInfo `json:"info"`
}
type CreateListResponse struct {
	ListId uint64 `json:"id"`
}

type UpdateListRequest struct {
	ListId   uint64         `json:"id"`
	ListInfo UpdateListInfo `json:"info"`
}
type DeleteListRequest struct {
	ListId uint64 `json:"id"`
}

type UpdateListResponse struct {
	Status string `json:"status"`
}

type DeleteListResponse struct {
	Status string `json:"status"`
}
