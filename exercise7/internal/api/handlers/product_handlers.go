package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/diegovanne/go23/exercise7/internal/api/entities"
	"github.com/diegovanne/go23/exercise7/internal/api/inputs"
	"github.com/diegovanne/go23/exercise7/internal/api/commons"
	"github.com/diegovanne/go23/exercise7/internal/api/database"
)

type ProductHandler struct {
}

var productList []entities.Product

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var input inputs.CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		commons.ResponseError(c, http.StatusBadRequest, "Invalid inputs", commons.ParseError(err))
		return
	}

	user := commons.GetUserAuth(c)

	product := entities.Product {
		Name:     input.Name,
		Price:    input.Price,
		Quantity: input.Quantity,
		UserID: user.ID,
	}

	if err := database.DB.Create(&product).Error; err != nil {
		commons.ResponseError(c, http.StatusInternalServerError, "Failed to create product", nil)
		return
	}

	result := struct {
		ID uint `json:"id"`
	}{
		ID: product.ID,
	}

	commons.ResponseSuccess(c, http.StatusCreated, "Product was created successfully", result.ID)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	productID := c.Param("product_id")
	var updatedProduct inputs.UpdateProductInput
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		commons.ResponseError(c, http.StatusBadRequest, "Invalid request payload", commons.ParseError(err))
		return
	}

	var product entities.Product
	if err := database.DB.Where("id = ?", productID).First(&product).Error; err != nil {
		commons.ResponseError(c, http.StatusNotFound, "Product not found", nil)
		return
	}

	product.Name = updatedProduct.Name
	product.Price = updatedProduct.Price
	product.Quantity = updatedProduct.Quantity

	if err := database.DB.Save(&product).Error; err != nil {
		commons.ResponseError(c, http.StatusInternalServerError, "Failed to update product", nil)
		return
	}

	c.JSON(http.StatusOK, product)
}


func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	productID := c.Param("product_id")

	var product entities.Product
	if err := database.DB.Where("id = ?", productID).First(&product).Error; err != nil {
		commons.ResponseError(c, http.StatusNotFound, "Product not found", nil)
		return
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		commons.ResponseError(c, http.StatusInternalServerError, "Failed to delete product", nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	user := commons.GetUserAuth(c)
	var products []entities.Product

	if err := database.DB.Where("created_by_id = ?", user.ID).Find(&products).Error; err != nil {
		commons.ResponseError(c, http.StatusInternalServerError, "Failed to fetch products", nil)
		return
	}

	c.JSON(http.StatusOK, products)
}
