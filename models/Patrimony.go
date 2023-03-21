package models

type Patrimony struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"Name"`
	Model string `json:"model"`
}
