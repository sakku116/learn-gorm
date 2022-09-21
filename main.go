package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	router := gin.Default()

	router.GET("products/", GetAllProducts)
	router.GET("products/:id/", GetProduct)
	router.PUT("products/:id/", UpdateProduct)
	router.POST("products/", CreateProduct)
	router.DELETE("products/:id/", DeleteProduct)

	router.Run()
}
