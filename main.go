package main

import (
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

	// Create products
	obs := "Super cool laptop"
	products := []model.Product{
		{
			Name:         "Laptop",
			Price:        1350,
			Observations: &obs,
		},
		{
			Name:  "Smartphone",
			Price: 350,
		},
		{
			Name:  "Monitor",
			Price: 250,
		},
	}

	storage.DB().CreateInBatches(products, len(products))
}
