package routes

import (
	"github.com/diegovanne/go23/exercise7/internal/api/handlers"
	"github.com/diegovanne/go23/exercise7/internal/api/middlewares"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	router.Use(middlewares.AuthMiddleware)

	productGroup := router.Group("/products")
	productGroup.Use(middlewares.JWTMiddleware())

	productHandler := handlers.NewProductHandler()

	productGroup.POST("/", productHandler.CreateProduct)
	productGroup.PUT("/:product_id", productHandler.UpdateProduct)
	productGroup.DELETE("/:product_id", productHandler.DeleteProduct)
	productGroup.GET("/products", productHandler.GetProducts)

	authGroup := router.Group("/auth")
	authHandler := handlers.NewAuthHandler()
	authGroup.POST("/login", authHandler.Login)

	cartGroup := router.Group("/cart")
	cartGroup.Use(middlewares.JWTMiddleware())
	cartHandler := handlers.NewCartHandler()
	cartGroup.POST("/add", cartHandler.AddToCart)
	cartGroup.POST("/remove", cartHandler.RemoveFromCart)
	cartGroup.POST("/checkout", cartHandler.Checkout)
}
