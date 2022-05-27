package models

type Album struct {
	ID     string  `json:"id,omitempty"`
	Title  string  `json:"title,omitempty"`
	Artist string  `json:"artist,omitempty"`
	Price  float64 `json:"price,omitempty"`
}
