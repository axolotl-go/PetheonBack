package servicestatushistory

import "gorm.io/gorm"

type ServiceStatusHistory struct {
	gorm.Model

	ServiceOrderID uint `gorm:"not null;index" json:"service_order_id"`

	Status string `gorm:"not null" json:"status"`
	Notes  string `json:"notes"`

	ChangedBy uint `gorm:"not null" json:"changed_by"`
}
