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

type HallController interface {
	FindAll(c *gin.Context) (int, interface{}, bool, error)
	FindByID(c *gin.Context) (int, interface{}, bool, error)
	Create(c *gin.Context) (int, interface{}, bool, error)
	Update(c *gin.Context) (int, interface{}, bool, error)
	Delete(c *gin.Context) (int, interface{}, bool, error)
}

type hallController struct {
	hallService service.HallService
}

func NewHallController(service service.HallService) HallController {
	return hallController{
		hallService: service,
	}
}

func (hc hallController) FindAll(c *gin.Context) (int, interface{}, bool, error) {
	halls, err := hc.hallService.FindAll()
	if err != nil {
		return http.StatusInternalServerError, nil, false, fmt.Errorf("could not get: %v", err)
	}

	return http.StatusOK, halls, true, nil
}

func (hc hallController) FindByID(c *gin.Context) (int, interface{}, bool, error) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		return http.StatusBadRequest, nil, false, errors.New("invalid query parameters values")
	}

	hall, err := hc.hallService.FindByID(id)
	if err != nil {
		return http.StatusInternalServerError, nil, false, fmt.Errorf("could not get: %v", err)
	}

	return http.StatusOK, hall, true, nil
}

func (hc hallController) Create(c *gin.Context) (int, interface{}, bool, error) {
	dataForNewHall := dto.HallCreate{}
	if err := c.ShouldBindJSON(&dataForNewHall); err != nil {
		return http.StatusBadRequest, nil, false, errors.New("invalid Body")
	}

	err := hc.hallService.Create(dataForNewHall)
	if err != nil {
		return http.StatusInternalServerError, nil, false, fmt.Errorf("could not create: %v", err)
	}

	return http.StatusOK, "Created", true, nil
}

func (hc hallController) Update(c *gin.Context) (int, interface{}, bool, error) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		return http.StatusBadRequest, nil, false, errors.New("invalid query parameters values")
	}

	dataToUpdateHall := dto.HallUpdate{}
	if err := c.ShouldBindJSON(&dataToUpdateHall); err != nil {
		return http.StatusBadRequest, nil, false, errors.New("invalid Body")
	}

	err = hc.hallService.Update(id, dataToUpdateHall)
	if err != nil {
		return http.StatusInternalServerError, nil, false, fmt.Errorf("could not update: %v", err)
	}

	return http.StatusOK, "Updated", true, nil

}

func (hc hallController) Delete(c *gin.Context) (int, interface{}, bool, error) {
	idRaw := c.Param("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		return http.StatusBadRequest, nil, false, errors.New("invalid query parameters values")
	}

	err = hc.hallService.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, nil, false, fmt.Errorf("could not delete: %v", err)
	}

	return http.StatusInternalServerError, "Deleted", true, nil
}
