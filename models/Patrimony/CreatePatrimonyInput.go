package models

type CreatePatrimonyInput struct {
	Name  string `json:"name" binding:"required"`
	Model string `json:"model" binding:"required"`
}
