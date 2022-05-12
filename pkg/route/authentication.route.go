package route

import (
	"chen/pkg/controller"
	"chen/pkg/repository"
	"chen/pkg/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthEndpoints(router *gin.Engine, db *gorm.DB, responseWrapper ResponseFormatterFunc) {
	authRepo := repository.NewAuthenticationRepository(db)
	jwtService := service.NewJWTService()
	authService := service.NewAuthenticationService(authRepo, jwtService)
	authController := controller.NewAuthenticationController(authService)

	auth := router.Group("/auth")
	auth.GET("/login", responseWrapper(authController.Login))
	auth.POST("/register", responseWrapper(authController.Register))
}
