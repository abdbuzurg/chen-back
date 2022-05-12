package controller

import (
	"chen/pkg/dto"
	"chen/pkg/service"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BranchController interface {
	FindByID(c *gin.Context) (int, interface{}, bool, error)
	FindAll(c *gin.Context) (int, interface{}, bool, error)
	Create(c *gin.Context) (int, interface{}, bool, error)
	Update(c *gin.Context) (int, interface{}, bool, error)
	Delete(c *gin.Context) (int, interface{}, bool, error)
}

type branchController struct {
	branchService service.BranchService
}

func NewBranchController(service service.BranchService) BranchController {
	return branchController{
		branchService: service,
	}
}

func (bc branchController) FindAll(c *gin.Context) (int, interface{}, bool, error) {
	branches, err := bc.branchService.FindAll()
	if err != nil {
		return http.StatusInternalServerError, nil, false, fmt.Errorf("could not find %v", err)
	}

	return http.StatusOK, branches, true, nil
}

func (bc branchController) FindByID(c *gin.Context) (int, interface{}, bool, error) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		return http.StatusBadRequest, nil, false, errors.New("invalid url parameters values")
	}

	branch, err := bc.branchService.FindByID(id)
	if err != nil {
		return http.StatusInternalServerError, nil, false, fmt.Errorf("could not get branches: %v", err)
	}

	return http.StatusOK, branch, true, nil
}

func (bc branchController) Create(c *gin.Context) (int, interface{}, bool, error) {
	dataForNewBranch := dto.BranchCreate{}
	if err := c.ShouldBindJSON(&dataForNewBranch); err != nil {
		return http.StatusBadRequest, nil, false, errors.New("invalid Body")
	}

	err := bc.branchService.Create(dataForNewBranch)
	if err != nil {
		return http.StatusInternalServerError, nil, false, fmt.Errorf("could not create: %v", err)
	}

	return http.StatusOK, "Created", true, nil
}

func (bc branchController) Update(c *gin.Context) (int, interface{}, bool, error) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		return http.StatusBadRequest, nil, false, errors.New("invalid query parameters values")
	}

	dataForNewBranch := dto.BranchUpdate{}
	if err := c.ShouldBindJSON(&dataForNewBranch); err != nil {
		return http.StatusBadRequest, nil, false, errors.New("invalid Body")
	}

	err = bc.branchService.Update(id, dataForNewBranch)
	if err != nil {
		return http.StatusInternalServerError, nil, false, fmt.Errorf("could not update: %v", err)
	}

	return http.StatusOK, "Updated", true, nil
}

func (bc branchController) Delete(c *gin.Context) (int, interface{}, bool, error) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		return http.StatusBadRequest, nil, false, errors.New("invalid query parameters values")
	}
	err = bc.branchService.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, nil, false, fmt.Errorf("could not delete: %v", err)
	}
	return http.StatusInternalServerError, "Deleted", true, nil
}
