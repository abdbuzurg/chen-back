package controller

import (
	"chen/pkg/dto"
	"chen/pkg/service"
	"chen/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HallController interface {
	Find(c *gin.Context)
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

func (hc hallController) Find(c *gin.Context) {
	idRaw := c.DefaultQuery("id", "0")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid query parameters values", false)
		return
	}

	halls, err := hc.hallService.Find(id)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Could not find", false)
		return
	}

	response.FormatResponse(c, http.StatusOK, halls, true)
}

func (hc hallController) Create(c *gin.Context) {
	dataForNewHall := dto.HallCreateDTO{}
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
	idRaw, exists := c.GetQuery("id")
	if !exists {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid query parameters", false)
		return
	}
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid Parameters", false)
		return
	}

	dataToUpdateHall := dto.HallUpdateDTO{}
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
	idRaw, exists := c.GetQuery("id")
	if !exists {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid query parameters", false)
		return
	}
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
