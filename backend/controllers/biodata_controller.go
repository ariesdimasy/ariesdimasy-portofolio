package controllers

import (
	"ariesdimasy-portofolio/models"
	"ariesdimasy-portofolio/services"
	"ariesdimasy-portofolio/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BiodataController interface {
	CreateBiodata(c *gin.Context)
	UpdateBiodata(c *gin.Context)
	DeleteBiodata(c *gin.Context)
	GetBiodataByID(c *gin.Context)
}

type biodataController struct {
	biodataService services.BiodataService
}

func NewBiodataController(biodataService services.BiodataService) BiodataController {
	return &biodataController{biodataService: biodataService}
}

func (bc biodataController) CreateBiodata(c *gin.Context) {
	var req models.BiodataCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	biodata := models.Biodata{
		UserID:   req.UserID,
		Headline: req.Headline,
		About:    req.About,
		Address:  req.Address,
		Phone:    req.Phone,
	}

	if err := bc.biodataService.CreateBiodata(&biodata); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Biodata created successfully"})
}

func (bc biodataController) UpdateBiodata(c *gin.Context) {
	id := c.Param("id")
	uintID, err := utils.StringToUint(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req models.BiodataUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	biodata, err := bc.biodataService.GetBiodataByID(uintID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Biodata not found"})
		return
	}

	biodata.Headline = req.Headline
	biodata.About = req.About
	biodata.Address = req.Address
	biodata.Phone = req.Phone

	if err := bc.biodataService.UpdateBiodata(biodata); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Biodata updated successfully"})
}

func (bc biodataController) DeleteBiodata(c *gin.Context) {
	var req models.BiodataDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	biodata, err := bc.biodataService.GetBiodataByID(req.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Biodata not found"})
		return
	}

	if err := bc.biodataService.DeleteBiodata(biodata); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Biodata deleted successfully"})
}

func (bc biodataController) GetBiodataByID(c *gin.Context) {
	id := c.Param("id")
	uintID, err := utils.StringToUint(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := bc.biodataService.GetBiodataByID(uintID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
