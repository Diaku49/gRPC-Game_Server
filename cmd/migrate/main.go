package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
)

func main() {
	db_URL := os.Getenv("DB_URL")

	err := migrateDB(db_URL)
	if err != nil {
		os.Exit(1)
	}
}

func migrateDB(URL string) error {
	m, err := migrate.New(
		"../../db/migrations",
		URL,
	)
	if err != nil {
		return fmt.Errorf("Couldnt make migrate instance :%v", err)
	}
	defer m.Close()

	if err = m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("Database is up to date.")
			return nil
		} else {
			log.Printf("Migration failed: %v", err)
			return fmt.Errorf("migration failed")
		}
	}

	return nil
}
