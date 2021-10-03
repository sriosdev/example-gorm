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

	// Retrive all products
	products := make([]model.Product, 0)
	storage.DB().Find(&products)

	for _, product := range products {
		fmt.Printf("%d - %s\n", product.ID, product.Name)
	}
}
