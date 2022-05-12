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

type TableController interface {
	FindAll(c *gin.Context) (int, interface{}, bool, error)
	FindByID(c *gin.Context) (int, interface{}, bool, error)
	Create(c *gin.Context) (int, interface{}, bool, error)
	Update(c *gin.Context) (int, interface{}, bool, error)
	Delete(c *gin.Context) (int, interface{}, bool, error)
}

type tableController struct {
	tableService service.TableService
}

func NewTableController(service service.TableService) TableController {
	return tableController{
		tableService: service,
	}
}

func (tc tableController) FindAll(c *gin.Context) (int, interface{}, bool, error) {
	tables, err := tc.tableService.FindAll()
	if err != nil {
		return http.StatusInternalServerError, nil, false, fmt.Errorf("could not get: %v", err)
	}

	return http.StatusOK, tables, true, nil
}

func (tc tableController) FindByID(c *gin.Context) (int, interface{}, bool, error) {
	idRaw := c.DefaultQuery("id", "0")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		return http.StatusBadRequest, nil, false, errors.New("invalid query parameters values")
	}

	table, err := tc.tableService.FindByID(id)
	if err != nil {
		return http.StatusInternalServerError, nil, false, fmt.Errorf("could not get: %v", err)
	}

	return http.StatusOK, table, true, nil
}

func (tc tableController) Create(c *gin.Context) (int, interface{}, bool, error) {
	dataForNewTable := dto.Table{}
	if err := c.ShouldBindJSON(&dataForNewTable); err != nil {
		return http.StatusBadRequest, nil, false, errors.New("invalid Body")
	}

	err := tc.tableService.Create(dataForNewTable)
	if err != nil {
		return http.StatusInternalServerError, nil, false, fmt.Errorf("could not create: %v", err)
	}

	return http.StatusOK, "Created", true, nil
}

func (tc tableController) Update(c *gin.Context) (int, interface{}, bool, error) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		return http.StatusBadRequest, nil, false, errors.New("invalid query parameters values")
	}

	dataToUpdateTable := dto.Table{}
	if err := c.ShouldBindJSON(&dataToUpdateTable); err != nil {
		return http.StatusBadRequest, nil, false, errors.New("invalid Body")
	}

	err = tc.tableService.Update(id, dataToUpdateTable)
	if err != nil {
		return http.StatusInternalServerError, nil, false, fmt.Errorf("could not update: %v", err)
	}

	return http.StatusOK, "Updated", true, nil
}

func (tc tableController) Delete(c *gin.Context) (int, interface{}, bool, error) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		return http.StatusBadRequest, nil, false, errors.New("invalid query parameters values")
	}

	err = tc.tableService.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, nil, false, fmt.Errorf("could not delete: %v", err)
	}

	return http.StatusInternalServerError, "Deleted", true, nil
}
