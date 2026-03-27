package serviceorders

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	OrderNumber string `gorm:"uniqueIndex;not null" json:"order_number"`

	UserID uint `gorm:"not null;index" json:"user_id"`
	PetID  uint `gorm:"not null;index" json:"pet_id"`

	ServiceType string `gorm:"not null;default:composta" json:"service_type"` // Composta & Cremation & Normal

	PickupRequired bool   `gorm:"not null" json:"pickup_required"`
	PickupAddress  string `json:"pickup_address"`

	UrnRequested bool `gorm:"not null" json:"urn_requested"` // Urna para las cenizas

	Price  float64 `gorm:"not null" json:"price"`
	Status string  `gorm:"not null;index" json:"status"`
}
