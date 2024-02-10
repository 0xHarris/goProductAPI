package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float32 `json:"price"`
	Stock       int     `json:"stock"`
}

var products = []product{
	{ID: "1", Name: "Big Clock", Description: "This is a Clock", Category: "Clocks", Price: 9.99, Stock: 6},
	{ID: "2", Name: "Stainless Steel Water Bottle", Description: "This is a Stainless Steel  Water Bottle", Category: "Kitchen", Price: 12.49, Stock: 1},
	{ID: "3", Name: "Wireless Headphones", Description: "Wireless Bluetooth Headphones", Category: "Clocks", Price: 99.99, Stock: 2},
}

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func productById(c *gin.Context) {
	id := c.Param("id")
	product, err := getProductById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not Found"})
		return
	}

	c.IndentedJSON(http.StatusOK, product)

}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")

	// Find the product index based on ID and remove it from the slice
	for i, p := range products {
		if p.ID == id {
			// Remove the product from the slice
			products = append(products[:i], products[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Product deleted"})
			return
		}
	}

	// If the product with the specified ID was not found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
}

func getProductById(id string) (*product, error) {
	for i, p := range products {
		if p.ID == id {
			return &products[i], nil
		}
	}

	return nil, errors.New("product not found")
}

func addProduct(c *gin.Context) {
	var newProduct product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)

}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.POST("/products", addProduct)
	router.GET("/products/:id", productById)
	router.DELETE("/products/:id", deleteProduct)
	router.Run("localhost:8080")
}
