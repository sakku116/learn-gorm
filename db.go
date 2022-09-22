package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB = connectDB()

func connectDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgres://postgres:root@localhost:5432/store"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var product_model Product
	db.AutoMigrate(&product_model)

	return db
}
