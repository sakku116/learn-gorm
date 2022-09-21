package main

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string
	Price int
}
type ProductInput struct {
	Name  string `json:"name" form:"name"`
	Price int    `json:"price" form:"price"`
}
