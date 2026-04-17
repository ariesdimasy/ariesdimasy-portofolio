package controllers

import (
	"ariesdimasy-portofolio/models"
	"ariesdimasy-portofolio/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BiodataController struct {
	BiodataService services.BiodataService
}

func NewBiodataController(biodataService services.BiodataService) *BiodataController {
	return &BiodataController{BiodataService: biodataService}
}

func (bc *BiodataController) CreateBiodata(c *gin.Context) {
	var biodata models.Biodata
	if err := c.ShouldBindJSON(&biodata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bc.BiodataService.CreateBiodata(&biodata); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Biodata created successfully"})
}

func (bc *BiodataController) UpdateBiodata(c *gin.Context) {
	var biodata models.Biodata
	if err := c.ShouldBindJSON(&biodata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bc.BiodataService.UpdateBiodata(&biodata); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Biodata updated successfully"})
}

func (bc *BiodataController) DeleteBiodata(c *gin.Context) {
	var biodata models.Biodata
	if err := c.ShouldBindJSON(&biodata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bc.BiodataService.DeleteBiodata(&biodata); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Biodata deleted successfully"})
}

func (bc *BiodataController) GetBiodataByID(c *gin.Context) {
	var biodata models.Biodata
	if err := c.ShouldBindJSON(&biodata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := bc.BiodataService.GetBiodataByID(biodata.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
