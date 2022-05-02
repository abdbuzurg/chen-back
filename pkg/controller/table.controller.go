package controller

import (
	"chen/pkg/dto"
	"chen/pkg/service"
	"chen/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TableController interface {
	Find(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type tableController struct {
	tableService service.TableService
}

func NewTableController(service service.TableService) TableController {
	return tableController{
		tableService: service,
	}
}

func (tc tableController) Find(c *gin.Context) {
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

	if id == 0 {
		tables, err := tc.tableService.FindAll()
		if err != nil {
			response.FormatResponse(c, http.StatusInternalServerError, "Could not get tables", false)
			return
		}

		response.FormatResponse(c, http.StatusOK, tables, true)
		return
	}

	table, err := tc.tableService.FindById(id)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Could not find", false)
		return
	}

	response.FormatResponse(c, http.StatusOK, table, true)
}

func (tc tableController) Create(c *gin.Context) {
	dataForNewTable := dto.TableDTO{}
	if err := c.ShouldBindJSON(&dataForNewTable); err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid body", false)
		return
	}

	err := tc.tableService.Create(dataForNewTable)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Could not create", false)
		return
	}

	response.FormatResponse(c, http.StatusOK, "Created", true)
}

func (tc tableController) Update(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid Parameters", false)
		return
	}

	dataToUpdateTable := dto.TableDTO{}
	if err := c.ShouldBindJSON(&dataToUpdateTable); err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid body", false)
		return
	}

	err = tc.tableService.Update(id, dataToUpdateTable)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Could not update", false)
		return
	}

	response.FormatResponse(c, http.StatusOK, "Updated", true)
}

func (tc tableController) Delete(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid Parameters", false)
		return
	}

	err = tc.tableService.Delete(id)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Could not delete", false)
		return
	}

	response.FormatResponse(c, http.StatusOK, "Deleted", true)
}
