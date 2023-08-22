package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/diegovanne/go23/exercise8/internal/api/inputs"
	"github.com/diegovanne/go23/exercise8/internal/api/entities"
	"github.com/diegovanne/go23/exercise8/internal/api/commons"
	"github.com/diegovanne/go23/exercise8/internal/api/database"
	"gorm.io/gorm"
)

type CartHandler struct{}

func NewCartHandler() *CartHandler {
	return &CartHandler{}
}

func (h *CartHandler) AddToCart(c *gin.Context) {
    var cartInsertInput inputs.CartInsertInput

    if err := c.ShouldBindJSON(&cartInsertInput); err != nil {
        commons.ResponseError(c, http.StatusBadRequest, "Invalid input", nil)
        return
    }

    currentUser := commons.GetUserAuth(c)

	var user entities.User
	if err := database.DB.Where("id = ?", currentUser.ID).First(&user).Error; err != nil {
		commons.ResponseError(c, http.StatusNotFound, "User not found", nil)
		return
	}

    var product entities.Product
    if err := database.DB.Where("id = ?", cartInsertInput.ProductID).First(&product).Error; err != nil {
        commons.ResponseError(c, http.StatusNotFound, "Product not found", nil)
        return
    }

    cart := entities.Cart{
        ProductID: cartInsertInput.ProductID,
        Name:      cartInsertInput.Name,
        Price:     cartInsertInput.Price,
        Quantity:  cartInsertInput.Quantity,
    }

    user.Products = append(user.Products, &product) 
    product.Quantity -= cartInsertInput.Quantity    

    if err := database.DB.Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(&cart).Error; err != nil {
            return err
        }
        if err := tx.Save(&user).Error; err != nil {
            return err
        }
        if err := tx.Save(&product).Error; err != nil {
            return err
        }
        return nil
    }); err != nil {
        commons.ResponseError(c, http.StatusInternalServerError, "Failed to create cart", nil)
        return
    }

    result := struct {
        ID uint `json:"id"`
    }{
        ID: cart.ID,
    }

    commons.ResponseSuccess(c, http.StatusCreated, "Cart was created successfully", result.ID)
}


func (h *CartHandler) RemoveFromCart(c *gin.Context) {
	cartID := c.Param("cart_id")

	var cart entities.Cart
	if err := database.DB.Where("id = ?", cartID).First(&cart).Error; err != nil {
		commons.ResponseError(c, http.StatusNotFound, "Cart not found", nil)
		return
	}

	if err := database.DB.Delete(&cart).Error; err != nil {
		commons.ResponseError(c, http.StatusInternalServerError, "Failed to delete cart", nil)
		return
	}

	commons.ResponseSuccess(c, http.StatusCreated, "Cart was deleted successfully", true)
}

func (h *CartHandler) Checkout(c *gin.Context) {
    currentUser := commons.GetUserAuth(c)

	var user entities.User
	if err := database.DB.Where("id = ?", currentUser.ID).First(&user).Error; err != nil {
		commons.ResponseError(c, http.StatusNotFound, "User not found", nil)
		return
	}

	var total float64

    if err := database.DB.Preload("Products").Model(&user).Association("Products").Find(&user.Products).Error; err != nil {
        commons.ResponseError(c, http.StatusInternalServerError, "Failed to get user's cart", nil)
        return
    }

    for _, product := range user.Products {
        total += product.Price * float64(product.Quantity)
    }

    commons.ResponseSuccess(c, http.StatusOK, "Checkout successful", gin.H{"total_price": total})
}
