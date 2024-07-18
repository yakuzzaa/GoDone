package serializer

type ItemInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type Item struct {
	ID       uint     `json:"id"`
	Info     ItemInfo `json:"info"`
	CreateAt string   `json:"create_at"`
	UpdateAt string   `json:"update_at"`
}
