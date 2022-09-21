package main

import "github.com/gin-gonic/gin"

func GetAllProducts(c *gin.Context) {
	var results []Product
	query := DB.Find(&results)
	if query.Error != nil {
		c.IndentedJSON(400, gin.H{
			"status": "error",
		})
		return
	}

	c.IndentedJSON(200, gin.H{
		"result": results,
	})
}

func CreateProduct(c *gin.Context) {
	// get input data from request
	var input ProductInput
	err := c.Bind(&input)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"status":  "failed",
			"message": "Bind error",
		})
		return
	}

	// create
	new_product := Product{
		Name:  input.Name,
		Price: input.Price,
	}
	query := DB.Create(&new_product)
	if query.Error != nil {
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

func GetProduct(c *gin.Context) {
	product_id := c.Param("id")

	// get product using primary key
	var result Product
	query := DB.Find(&result, product_id)
	if query.Error != nil {
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

func UpdateProduct(c *gin.Context) {
	product_id := c.Param("id")

	// get input from request
	var input ProductInput
	err := c.Bind(&input)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"status": "input data required",
		})
		return
	}

	// update
	var result Product
	/* using "Updates" instead "Update" because we dont know what the fields that user want to update
	using "Where" to get and upadte at one tiem */
	query := DB.Model(&result).Where("id = ?", product_id).Updates(Product{
		Name:  input.Name,
		Price: input.Price,
	})
	if query.Error != nil {
		c.IndentedJSON(400, gin.H{
			"status": "failed",
		})
		return
	}

	// response
	c.IndentedJSON(200, gin.H{
		"status": "success",
		"data":   result,
	})
}

func DeleteProduct(c *gin.Context) {
	product_id := c.Param("id")

	// delete product by primary key
	var product Product
	/* use "Unscoped" to permanently delete the record instead just filling "deleted_at" field (soft delete) */
	query := DB.Unscoped().Delete(&product, product_id)
	if query.Error != nil {
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
