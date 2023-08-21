package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type CartHandler struct{}

func NewCartHandler() *CartHandler {
	return &CartHandler{}
}

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

var cart []Product

func AddToCart(c *gin.Context) {
	var input struct {
		ProductID int `json:"product_id"`
		Quantity  int `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	product := Product{
		ID:       input.ProductID,
		Name:     "Sample Product",
		Price:    100,
		Quantity: input.Quantity,
	}

	cart = append(cart, product)

	c.JSON(http.StatusOK, gin.H{"message": "Product added to cart"})
}

func RemoveFromCart(c *gin.Context) {
	var input struct {
		ProductID int `json:"product_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	for i, item := range cart {
		if item.ID == input.ProductID {
			cart = append(cart[:i], cart[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Product removed from cart"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found in cart"})
}

func Checkout(c *gin.Context) {
	totalPrice := 0

	for _, item := range cart {
		totalPrice += item.Price * item.Quantity
	}

	cart = nil

	c.JSON(http.StatusOK, gin.H{"message": "Checkout successful", "total_price": totalPrice})
}