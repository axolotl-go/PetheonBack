package servicetype

import "gorm.io/gorm"

type ServiceType struct {
	gorm.Model

	name        string
	description string
	price       float64
}
