package model

import "gorm.io/gorm"

// InvoiceItem model
type InvoiceItem struct {
	gorm.Model
	InvoiceHeaderID uint
	ProductID       uint
}
