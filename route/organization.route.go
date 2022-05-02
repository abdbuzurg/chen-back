package route

import (
	"chen/pkg/controller"
	"chen/pkg/middleware"
	"chen/pkg/repository"
	"chen/pkg/service"
)

func (s *Server) OrganizationCRUDEndpoints() {
	organizationRepo := repository.NewOrganizationRepo(s.DB)
	organizationService := service.NewOrganizationService(organizationRepo)
	organizationController := controller.NewOrganizationController(organizationService)

	org := s.Router.Group("/organization", middleware.AuthMiddleware())
	org.GET("", organizationController.Find)
	org.POST("", organizationController.Create)
	org.PUT("/:id", organizationController.Update)
	org.DELETE("/:id", organizationController.Delete)
}
