package models

type Student struct {
	ID     uint    `json:"id" gorm:"primary_key"`
	Name   string  `json:"name"`
	Grade  uint8   `json:"grade"`
	Rating float32 `json:"rating"`
}
