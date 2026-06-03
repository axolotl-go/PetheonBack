package formregister

import (
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model

	Name           string  `json:"name" validate:"required,min=2,max=100"`
	Phone          string  `json:"phone" validate:"required,min=8,max=20"`
	Service        string  `json:"service" validate:"required"`
	Species        string  `json:"species" validate:"required"`
	PetName        string  `json:"pet_name"`
	ApproxWeightKg float64 `json:"approx_weight_kg" validate:"required,gt=0"`
}
