package serviceorders

import (
	"github.com/axolotl-go/eternal_paw/internal/pets"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	UserID uint `json:"user_id"`

	PetID uint     `json:"pet_id"`
	Pet   pets.Pet `json:"pet" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	OrderNumber   string `json:"order_number"`
	ServiceTypeID uint   `json:"service_type_id" validate:"required"`

	PickupRequired bool   `json:"pickup_required"`
	PickupAddress  string `json:"pickup_address" validate:"required_if=PickupRequired true"`

	Active bool    `json:"active" gorm:"default:false"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

type OrderResponse struct {
}
