package main

import "github.com/gin-gonic/gin"

func GetAllProducts(c *gin.Context) {
	var results []Product // create list of type to get some product
	get := DB.Find(&results)
	if get.Error != nil {
		c.IndentedJSON(400, gin.H{
			"status": "error",
		})
		return
	}

	// final response
	c.IndentedJSON(200, gin.H{
		"result": results,
	})
}

func GetProduct(c *gin.Context) {
	product_id := c.Param("id")

	// get product using primary key
	var result Product
	get := DB.Find(&result, product_id)
	if get.Error != nil {
		c.IndentedJSON(400, gin.H{
			"status": "failed",
		})
		return
	}

	// response
	c.IndentedJSON(200, gin.H{
		"result": result,
	})
}

func CreateProduct(c *gin.Context) {
	// get input from request
	var new_product Product
	err := c.Bind(&new_product)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"status":  "failed",
			"message": "Bind error",
		})
		return
	}

	// create
	create := DB.Create(&new_product)
	if create.Error != nil {
		c.IndentedJSON(400, gin.H{
			"status": "failed",
		})
		return
	}

	// response
	c.IndentedJSON(200, gin.H{
		"status": "success",
	})
}

func UpdateProduct(c *gin.Context) {
	product_id := c.Param("id")

	// get input from request
	var new_data Product
	err := c.Bind(&new_data)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"status": "input data required",
		})
		return
	}

	// get
	var result Product
	get := DB.Find(&result)
	if get.Error != nil {
		c.IndentedJSON(400, gin.H{
			"status": "failed",
		})
		return
	}

	/*
		use ".Model(&[returned get query])"  to get actual record and then update it,
		so you dont need to make find query again to get actual data for the response!

		but u still can remove .Model() if you dont mind about returned actual data in response
	*/

	// update
	update := DB.Model(&result).Where("id = ?", product_id).Updates(&new_data)
	if update.Error != nil {
		c.IndentedJSON(400, gin.H{
			"status": "failed",
		})
		return
	}

	// final response
	c.IndentedJSON(200, gin.H{
		"status": "success",
		"data":   result,
	})
}

func DeleteProduct(c *gin.Context) {
	product_id := c.Param("id")

	// delete
	var product Product
	/* use "Unscoped" to permanently delete the record instead just filling "deleted_at" field (soft delete) */
	query := DB.Unscoped().Delete(&product, product_id)
	/* or can use "Where" to specify confition */
	if query.Error != nil {
		c.IndentedJSON(400, gin.H{
			"status": "failed",
		})
		return
	}

	// final response
	c.IndentedJSON(200, gin.H{
		"status": "success",
	})
}

func Playground(c *gin.Context) {
	product_id := c.Param("id")

	// query
	var product Product
	query := DB.Unscoped().Where("id = ?", product_id).Delete(&product)
	if query.Error != nil {
		c.IndentedJSON(400, gin.H{
			"message": "gorm error",
		})
		return
	}

	c.IndentedJSON(200, map[string]interface{}{
		"message": "success",
	})
}
