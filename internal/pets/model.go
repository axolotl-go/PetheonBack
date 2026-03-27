package pets

import (
	"time"

	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model

	UserID    uint       `gorm:"index;not null" json:"user_id"`
	ImageUrl  *string    `json:"image_url,omitempty"`
	Name      string     `gorm:"not null" json:"name"`
	Species   string     `gorm:"not null" json:"species"`
	Breed     string     `gorm:"not null" json:"breed"`
	Weight    float64    `gorm:"not null" json:"weight"`
	DeathDate *time.Time `json:"death_date,omitempty"`
}
