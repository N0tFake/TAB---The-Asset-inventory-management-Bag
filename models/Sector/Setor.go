package models

type Sector struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
}
