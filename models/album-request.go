package models

type AlbumRequest struct {
	Title  string  `json:"title,omitempty"`
	Artist string  `json:"artist,omitempty"`
	Price  float64 `json:"price,omitempty"`
}
