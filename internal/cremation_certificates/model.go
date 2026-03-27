package cremationcertificates

import "gorm.io/gorm"

type CremationCertificate struct {
	gorm.Model

	ServiceOrderID uint `gorm:"not null;uniqueIndex" json:"service_order_id"`

	CertificateURL string `gorm:"not null" json:"certificate_url"`
}
