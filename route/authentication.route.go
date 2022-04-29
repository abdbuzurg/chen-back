package route

import (
	"chen/pkg/controller"
	"chen/pkg/repository"
	"chen/pkg/service"
)

func (s *Server) AuthEndpoints() {
	authRepo := repository.NewAuthenticationRepository(s.DB)
	authService := service.NewAuthenticationService(authRepo)
	authController := controller.NewAuthenticationController(authService)

	auth := s.Router.Group("/auth")
	auth.GET("/login", authController.Login)
	auth.POST("/register", authController.Register)
}
