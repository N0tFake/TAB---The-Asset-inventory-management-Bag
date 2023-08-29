package models

import "gorm.io/gorm"

type Patrimony struct {
	gorm.Model
	SerialNumber string
	Name         string
	ModelName    string
	// Category          int
	SectorRef         int
	DateOfAcquisition string
	Age               int
	PurchasePrice     float32
	Lifespan          int
	CurrentValue      float64
	// ConservationState int     `json:"conservationState" gorm:"foreignKey:ID"`
}

type Sector struct {
	gorm.Model
	Name      string
	Tag       string
	Patrimony []Patrimony `gorm:"foreignKey:SectorRef"`
}
