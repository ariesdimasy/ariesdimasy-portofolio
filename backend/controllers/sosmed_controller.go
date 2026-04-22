package controllers

import (
	"ariesdimasy-portofolio/models"
	"ariesdimasy-portofolio/services"
	"ariesdimasy-portofolio/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SosmedController interface {
	CreateSosmed(c *gin.Context)
	UpdateSosmed(c *gin.Context)
	DeleteSosmed(c *gin.Context)
	GetSosmedByID(c *gin.Context)
	GetAllSosmeds(c *gin.Context)
}

type sosmedController struct {
	sosmedService services.SosmedService
}

func NewSosmedController(sosmedService services.SosmedService) SosmedController {
	return &sosmedController{sosmedService: sosmedService}
}

func (sc sosmedController) CreateSosmed(c *gin.Context) {
	var req models.SosmedCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sosmed := models.Sosmed{
		UserID: req.UserID,
		Name:   req.Name,
		Icon:   req.Icon,
		Link:   req.Link,
	}

	if err := sc.sosmedService.CreateSosmed(&sosmed); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sosmed created successfully"})
}

func (sc sosmedController) UpdateSosmed(c *gin.Context) {
	id := c.Param("id")
	uintID, err := utils.StringToUint(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req models.SosmedUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sosmed, err := sc.sosmedService.GetSosmedByID(uintID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sosmed not found"})
		return
	}

	sosmed.Name = req.Name
	sosmed.Icon = req.Icon
	sosmed.Link = req.Link

	if err := sc.sosmedService.UpdateSosmed(sosmed); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sosmed updated successfully"})
}

func (sc sosmedController) DeleteSosmed(c *gin.Context) {
	var req models.SosmedDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sosmed, err := sc.sosmedService.GetSosmedByID(req.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sosmed not found"})
		return
	}

	if err := sc.sosmedService.DeleteSosmed(sosmed); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sosmed deleted successfully"})
}

func (sc sosmedController) GetSosmedByID(c *gin.Context) {
	id := c.Param("id")
	uintID, err := utils.StringToUint(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sosmed, err := sc.sosmedService.GetSosmedByID(uintID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sosmed)
}

func (sc sosmedController) GetAllSosmeds(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	query := models.SosmedQuery{
		Name: c.Query("name"),
	}

	sosmeds, err := sc.sosmedService.GetAllSosmeds(userID, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sosmeds})
}
