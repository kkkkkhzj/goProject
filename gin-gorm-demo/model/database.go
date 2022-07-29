package model

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	engine *gorm.DB
}

// global variable
var currentDB *Database

func GetCurrentDB() *Database {
	if currentDB == nil {
		log.Panic("database is nil")
	}
	return currentDB
}

func SetupDatabase(sqlitePath string) (result *Database, err error) {
	db, err := gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema, create users table
	db.AutoMigrate(&User{})
	result = &Database{engine: db}
	currentDB = result
	return result, nil
}
