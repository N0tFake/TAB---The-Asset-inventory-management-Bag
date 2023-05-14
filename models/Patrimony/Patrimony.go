package models

type Patrimony struct {
	ID                uint    `json:"id" gorm:"primary_key"`
	SerialNumber      string  `json:"serialNumber"`
	Name              string  `json:"name"`
	Model             string  `json:"model"`
	Category          int     `json:"category" gorm:"foreignKey:CategoryRefer"`
	Sector            int     `json:"sector" gorm:"forignKey:SetorRefer"`
	DateOfAcquisition string  `json:"dateOfAcquisition"`
	Age               int     `json:"age"`
	PurchasePrice     float32 `json:"purchasePrice"`
	Lifespan          int     `json:"lifespan"`
	CurrentValue      float64 `json:"currentValue"`
	ConservationState int     `json:"conservationState"`
}
