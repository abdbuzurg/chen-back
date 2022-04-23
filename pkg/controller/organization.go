package controller

import (
	"chen/model"
	"chen/pkg/service"
	"chen/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OrgFind(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid parameters", false)
		return
	}
	org, err := service.OrgFind(id)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, err.Error(), false)
	}

	response.FormatResponse(c, http.StatusOK, org, true)
}
func OrgCreate(c *gin.Context) {
	var dataForCreatingNewOrg model.OrganizationData
	if err := c.BindJSON(&dataForCreatingNewOrg); err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid Body", false)
		return
	}

	err := service.OrgCreate(dataForCreatingNewOrg)

	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Could not create ORG", false)
		return
	}

	response.FormatResponse(c, http.StatusOK, "New Org is created", true)
}
func OrgUpdate(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid parameters", false)
		return
	}

	var dataForUpdaingOrgInfo model.OrganizationData
	if err := c.BindJSON(&dataForUpdaingOrgInfo); err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Invalid Body", false)
		return
	}

	updatedOrgInfo, err := service.OrgUpdate(id, dataForUpdaingOrgInfo)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Cannot update Org Info", false)
		return
	}

	response.FormatResponse(c, http.StatusOK, updatedOrgInfo, true)
}
func OrgDelete(c *gin.Context) {

	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid parameters", false)
		return
	}

	err = service.OrgDelete(id)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Could not delete org", false)
	}

	response.FormatResponse(c, http.StatusOK, "Successfully Deleted", true)
}