package controller

import (
	"chen/model"
	"chen/pkg/service"
	"chen/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HallController interface {
	FindAll(c *gin.Context)
	FindById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type hallController struct {
	hallService service.HallService
}

func NewHallController(service service.HallService) HallController {
	return hallController{
		hallService: service,
	}
}

func (hc hallController) FindAll(c *gin.Context) {
	halls, err := hc.hallService.FindAll()
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Could not get branches", false)
		return
	}

	response.FormatResponse(c, http.StatusOK, halls, true)
}

func (hc hallController) FindById(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid Parameters", false)
		return
	}

	hall, err := hc.hallService.FindById(id)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Could not find", false)
		return
	}

	response.FormatResponse(c, http.StatusOK, hall, true)
}

func (hc hallController) Create(c *gin.Context) {
	dataForNewHall := model.HallCreateData{}
	if err := c.ShouldBindJSON(&dataForNewHall); err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid body", false)
		return
	}

	err := hc.hallService.Create(dataForNewHall)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Could not create", false)
		return
	}

	response.FormatResponse(c, http.StatusOK, "Created", true)
}

func (hc hallController) Update(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid Parameters", false)
		return
	}

	dataToUpdateHall := model.HallUpdateData{}
	if err := c.ShouldBindJSON(&dataToUpdateHall); err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid body", false)
		return
	}

	err = hc.hallService.Update(id, dataToUpdateHall)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Could not update", false)
		return
	}

	response.FormatResponse(c, http.StatusOK, "Updated", true)
}

func (hc hallController) Delete(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid Parameters", false)
		return
	}

	err = hc.hallService.Delete(id)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Could not delete", false)
		return
	}

	response.FormatResponse(c, http.StatusOK, "Deleted", true)
}
