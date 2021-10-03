package main

import (
	"fmt"
	"log"

	"github.com/sriosdev/gorm-example/model"
	"github.com/sriosdev/gorm-example/storage"
)

func main() {
	drive := storage.Postgres
	storage.New(drive)

	// Migrate tables if it's necesary
	err := storage.DB().AutoMigrate(
		&model.Product{},
		&model.InvoiceHeader{},
		&model.InvoiceItem{},
	)
	if err != nil {
		log.Fatalf("Error happends during migration: %v\n", err)
	}

	// Retrive product by ID
	product := model.Product{}
	storage.DB().First(&product, 2)
	fmt.Printf("%d - %s\n", product.ID, product.Name)

	// Soft delete previous product
	storage.DB().Delete(&product)
}
