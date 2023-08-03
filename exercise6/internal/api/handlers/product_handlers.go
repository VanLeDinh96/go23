package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/diegovanne/go23/exercise6/internal/api/models"
)

type ProductHandler struct {
}

var productList []models.Product

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	product.ID = len(productList) + 1
	productList = append(productList, product)
	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	productID := c.Param("product_id")
	var updatedProduct models.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	id := 0
	for i, p := range productList {
		if fmt.Sprint(p.ID) == productID {
			id = i
			break
		}
	}

	if id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	productList[id].Name = updatedProduct.Name
	productList[id].Description = updatedProduct.Description
	productList[id].Price = updatedProduct.Price

	c.JSON(http.StatusOK, productList[id])
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	productID := c.Param("product_id")

	id := 0
	for i, p := range productList {
		if fmt.Sprint(p.ID) == productID {
			id = i
			break
		}
	}

	if id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	productList = append(productList[:id], productList[id+1:]...)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, productList)
}