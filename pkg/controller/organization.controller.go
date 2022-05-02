package controller

import (
	"chen/pkg/dto"
	"chen/pkg/service"
	"chen/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrganizationController interface {
	Find(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type organizationController struct {
	organizationService service.OrganizationService
}

func NewOrganizationController(service service.OrganizationService) OrganizationController {
	return organizationController{
		organizationService: service,
	}
}

func (oc organizationController) FindAll(c *gin.Context) {

}

func (oc organizationController) Find(c *gin.Context) {
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
		orgs, err := oc.organizationService.FindAll()
		if err != nil {
			response.FormatResponse(c, http.StatusInternalServerError, "could not fetch entries", false)
			return
		}

		response.FormatResponse(c, http.StatusOK, orgs, true)
		return
	}

	org, err := oc.organizationService.FindById(id)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, err.Error(), false)
		return
	}

	response.FormatResponse(c, http.StatusOK, org, true)
}

func (oc organizationController) Create(c *gin.Context) {
	var dataForCreatingNewOrg dto.OrganizationDTO
	if err := c.ShouldBindJSON(&dataForCreatingNewOrg); err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid Body", false)
		return
	}
	err := oc.organizationService.Create(dataForCreatingNewOrg)

	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Could not create ORG", false)
		return
	}

	response.FormatResponse(c, http.StatusOK, "New Org is created", true)
}

func (oc organizationController) Update(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid parameters", false)
		return
	}

	var dataForUpdaingOrgInfo dto.OrganizationDTO
	if err := c.ShouldBindJSON(&dataForUpdaingOrgInfo); err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid Body", false)
		return
	}
	err = oc.organizationService.Update(id, dataForUpdaingOrgInfo)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Cannot update Org Info", false)
		return
	}

	response.FormatResponse(c, http.StatusOK, "Info updated successfully", true)
}
func (oc organizationController) Delete(c *gin.Context) {

	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid parameters", false)
		return
	}

	err = oc.organizationService.Delete(id)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Could not delete org", false)
	}

	response.FormatResponse(c, http.StatusOK, "Successfully Deleted", true)
}
