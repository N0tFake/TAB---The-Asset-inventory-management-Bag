package models

import (
	category "tab/models/Category"
	conservation "tab/models/Conservation"
	patrimony "tab/models/Patrimony"
)

var (
	Patrimony    patrimony.Patrimony
	Category     category.Category
	Sector       patrimony.Sector
	Conservation conservation.Conservation
)
