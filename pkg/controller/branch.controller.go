package controller

import (
	"chen/pkg/dto"
	"chen/pkg/service"
	"chen/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BranchController interface {
	Find(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type branchController struct {
	branchService service.BranchService
}

func NewBranchController(service service.BranchService) BranchController {
	return branchController{
		branchService: service,
	}
}

func (bc branchController) Find(c *gin.Context) {
	idRaw, exists := c.GetQuery("id")
	if !exists {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid query parameters", false)
		return
	}

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid query parameters values", false)
		return
	}

	if id == 0 {
		// http://web.site/branch
		branches, err := bc.branchService.FindAll()
		if err != nil {
			response.FormatResponse(c, http.StatusInternalServerError, "Could not get branches", false)
			return
		}

		response.FormatResponse(c, http.StatusOK, branches, true)
		return
	}
	// http://web.site/branch?id=1
	branch, err := bc.branchService.FindById(id)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Could not find Branch", false)
		return
	}

	response.FormatResponse(c, http.StatusOK, branch, true)

}

func (bc branchController) Create(c *gin.Context) {
	dataForNewBranch := dto.BranchCreateDTO{}
	if err := c.ShouldBindJSON(&dataForNewBranch); err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid Body", false)
		return
	}

	err := bc.branchService.Create(dataForNewBranch)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Could not create", false)
		return
	}

	response.FormatResponse(c, http.StatusOK, "Created", true)
}

func (bc branchController) Update(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid Parameters", false)
		return
	}

	dataForNewBranch := dto.BranchUpdateDTO{}
	if err := c.ShouldBindJSON(&dataForNewBranch); err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid Body", false)
		return
	}

	err = bc.branchService.Update(id, dataForNewBranch)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Could not update", false)
		return
	}

	response.FormatResponse(c, http.StatusOK, "Updated", true)
}

func (bc branchController) Delete(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid Parameters", false)
		return
	}

	err = bc.branchService.Delete(id)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Could not delete", false)
		return
	}
	response.FormatResponse(c, http.StatusOK, "Deleted", true)
}
