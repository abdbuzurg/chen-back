package routes

import (
	"chen/pkg/controller"
	"chen/pkg/repository"
	"chen/pkg/service"
)

func (s *Server) BranchCRUDEndpoints() {
	branchRepo := repository.NewBranchRepository(s.DB)
	branchService := service.NewBranchService(branchRepo)
	branchController := controller.NewBranchController(branchService)

	branch := s.Router.Group("/branch")
	branch.GET("", branchController.FindAll)
	branch.GET("/:id", branchController.FindById)
	branch.POST("", branchController.Create)
	branch.PUT("/:id", branchController.Update)
	branch.DELETE("/:id", branchController.Delete)
}
