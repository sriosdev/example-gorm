package model

import (
	"gorm.io/gorm"
)

// InvoiceHeader model
type InvoiceHeader struct {
	gorm.Model
	Client       string `gorm:"type:varchar(100); not null"`
	InvoiceItems []InvoiceItem
}
