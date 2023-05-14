package models

type Conservation struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}
