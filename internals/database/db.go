package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Client() *gorm.DB {
	return DB
}

func Connect() {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		dsn = "user=manager database=postgres sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Unable to open database, error: %v", err)
		panic(err)
	}

	sql, err := db.DB()
	if err != nil {
		fmt.Printf("Unable to get sql database from gorm, err: %v", err)
		panic(err)
	}

	if err := sql.Ping(); err != nil {
		fmt.Printf("Unable to connect to the database, error: %v", err)
		panic(err)
	}

	fmt.Println("Successfully connected to the postgres db")

	DB = db
}
