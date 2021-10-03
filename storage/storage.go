package storage

import (
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Storage driver
type Driver string

const (
	Mariadb  Driver = "MARIADB"
	Postgres Driver = "POSTGRES"
)

// New creates the connection with the database
func New(d Driver) {
	switch d {
	case Mariadb:
		newMariaDB()
	case Postgres:
		newPostgresDB()
	}
}

func newPostgresDB() {
	var err error

	once.Do(func() {
		dsn := "postgres://gopher:secret@localhost:7530/gorm-example?sslmode=disable"
		db, err = gorm.Open(postgres.Open(dsn))
		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}

		log.Println("Successfuly connected to Postgres")
	})
}

func newMariaDB() {
	var err error

	once.Do(func() {
		dsn := "gopher:secret@tcp(localhost:7531)/gorm-example"
		db, err = gorm.Open(mysql.Open(dsn))
		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}

		log.Println("Successfuly connected to Mariadb")
	})
}

// Return a single instance of db
func DB() *gorm.DB {
	return db
}
