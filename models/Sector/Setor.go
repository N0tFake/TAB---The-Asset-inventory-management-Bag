package models

type Sector struct {
	ID   uint   `json:"sectorId" gorm:"primary_key"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
}
