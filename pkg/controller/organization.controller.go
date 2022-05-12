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

type OrganizationController interface {
	FindAll(c *gin.Context) (int, interface{}, bool, error)
	FindByID(c *gin.Context) (int, interface{}, bool, error)
	Create(c *gin.Context) (int, interface{}, bool, error)
	Update(c *gin.Context) (int, interface{}, bool, error)
	Delete(c *gin.Context) (int, interface{}, bool, error)
}

type organizationController struct {
	organizationService service.OrganizationService
}

func NewOrganizationController(service service.OrganizationService) OrganizationController {
	return organizationController{
		organizationService: service,
	}
}

func (oc organizationController) FindAll(c *gin.Context) (int, interface{}, bool, error) {
	organizations, err := oc.organizationService.FindAll()
	if err != nil {
		return http.StatusInternalServerError, nil, false, fmt.Errorf("could not get: %v", err)
	}

	return http.StatusOK, organizations, true, nil
}

func (oc organizationController) FindByID(c *gin.Context) (int, interface{}, bool, error) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		return http.StatusBadRequest, nil, false, errors.New("invalid url parameters values")
	}

	org, err := oc.organizationService.FindByID(id)
	if err != nil {
		return http.StatusInternalServerError, nil, false, fmt.Errorf("could not get branches: %v", err)
	}

	return http.StatusOK, org, true, nil
}

func (oc organizationController) Create(c *gin.Context) (int, interface{}, bool, error) {
	var dataForCreatingNewOrg dto.Organization
	if err := c.ShouldBindJSON(&dataForCreatingNewOrg); err != nil {
		return http.StatusBadRequest, nil, false, errors.New("invalid Body")
	}
	err := oc.organizationService.Create(dataForCreatingNewOrg)

	if err != nil {
		return http.StatusInternalServerError, nil, false, fmt.Errorf("could not create: %v", err)
	}

	return http.StatusOK, "Created", true, nil
}

func (oc organizationController) Update(c *gin.Context) (int, interface{}, bool, error) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		return http.StatusBadRequest, nil, false, errors.New("invalid query parameters values")
	}

	var dataForUpdaingOrgInfo dto.Organization
	if err := c.ShouldBindJSON(&dataForUpdaingOrgInfo); err != nil {
		return http.StatusBadRequest, nil, false, errors.New("invalid Body")
	}

	err = oc.organizationService.Update(id, dataForUpdaingOrgInfo)
	if err != nil {
		return http.StatusInternalServerError, nil, false, fmt.Errorf("could not update: %v", err)
	}

	return http.StatusOK, "Updated", true, nil
}
func (oc organizationController) Delete(c *gin.Context) (int, interface{}, bool, error) {
	idRaw := c.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		return http.StatusBadRequest, nil, false, errors.New("invalid query parameters values")
	}

	err = oc.organizationService.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, nil, false, fmt.Errorf("could not delete: %v", err)
	}

	return http.StatusInternalServerError, "Deleted", true, nil
}
