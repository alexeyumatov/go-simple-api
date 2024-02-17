package models

type Student struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Grade  uint8   `json:"grade"`
	Rating float32 `json:"rating"`
}
