package database

import (
	"log"

	"go.uber.org/dig"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init(c *dig.Container) {
	if err := c.Provide(initDB); err != nil {
		log.Fatalf("provide DB failed: %s", err.Error())
	}
}

func initDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("please.db"), &gorm.Config{})
}
