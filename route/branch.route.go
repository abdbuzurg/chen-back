package route

import (
	"chen/pkg/controller"
	"chen/pkg/middleware"
	"chen/pkg/repository"
	"chen/pkg/service"
)

func (s *Server) BranchCRUDEndpoints() {
	branchRepo := repository.NewBranchRepository(s.DB)
	branchService := service.NewBranchService(branchRepo)
	branchController := controller.NewBranchController(branchService)

	branch := s.Router.Group("/branch", middleware.AuthMiddleware())
	branch.GET("", branchController.Find)
	branch.POST("", branchController.Create)
	branch.PUT("", branchController.Update)
	branch.DELETE("", branchController.Delete)
}
