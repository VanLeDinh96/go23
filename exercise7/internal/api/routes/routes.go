package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/diegovanne/go23/exercise7/internal/api/handlers"
	"github.com/diegovanne/go23/exercise7/internal/api/middlewares"
)

func InitializeRoutes(router *gin.Engine) {
	router.Use(gin.Recovery())
	
	router.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "ok",
        })
    })

	authHandler := handlers.NewAuthHandler()
	productHandler := handlers.NewProductHandler()
	cartHandler := handlers.NewCartHandler()
	
	authGroup := router.Group("/auth")
	authGroup.POST("/login", authHandler.Login)
	authGroup.POST("/register", authHandler.Register)


	backOfficeGroup := router.Group("/")
	backOfficeGroup.Use(middlewares.JWTMiddleware())
	backOfficeGroup.Use(middlewares.BlacklistMiddleware())

	backOfficeGroup.POST("/logout", authHandler.Logout)
	backOfficeGroup.POST("/change-password", authHandler.ChangePassword)
	backOfficeGroup.GET("/me", authHandler.GetProfile)
	backOfficeGroup.PUT("/me", authHandler.UpdateProfile)
	backOfficeGroup.POST("/validate-token", authHandler.ValidateToken)

	productGroup := backOfficeGroup.Group("/products")
	{
		productGroup.POST("/", productHandler.CreateProduct)
		productGroup.PUT("/:product_id", productHandler.UpdateProduct)
		productGroup.DELETE("/:product_id", productHandler.DeleteProduct)
		productGroup.GET("/products", productHandler.GetProducts)
	}

	cartGroup := backOfficeGroup.Group("/cart")
	{
		cartGroup.POST("/add", cartHandler.AddToCart)
		cartGroup.POST("/remove", cartHandler.RemoveFromCart)
		cartGroup.POST("/checkout", cartHandler.Checkout)
	}
}
