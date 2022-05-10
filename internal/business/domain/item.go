package domain

type Item struct {
	Id        string  `json:"id,omitempty"`
	Title     string  `json:"title,omitempty"`
	Price     float32 `json:"price,omitempty"`
	Thumbnail string  `json:"thumbnail,omitempty"`
}
