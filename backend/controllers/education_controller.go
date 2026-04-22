package controllers

import (
	"ariesdimasy-portofolio/models"
	"ariesdimasy-portofolio/services"
	"ariesdimasy-portofolio/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EducationController interface {
	CreateEducation(c *gin.Context)
	UpdateEducation(c *gin.Context)
	DeleteEducation(c *gin.Context)
	GetEducationByID(c *gin.Context)
	GetAllEducations(c *gin.Context)
}

type educationController struct {
	educationService services.EducationService
}

func NewEducationController(educationService services.EducationService) EducationController {
	return &educationController{educationService: educationService}
}

func (ec educationController) CreateEducation(c *gin.Context) {
	var req models.EducationCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	education := models.Education{
		UserID:      req.UserID,
		Degree:      req.Degree,
		Major:       req.Major,
		Institution: req.Institution,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
	}

	if err := ec.educationService.CreateEducation(&education); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Education created successfully"})
}

func (ec educationController) UpdateEducation(c *gin.Context) {
	id := c.Param("id")
	uintID, err := utils.StringToUint(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req models.EducationUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	education, err := ec.educationService.GetEducationByID(uintID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Education not found"})
		return
	}

	education.Degree = req.Degree
	education.Major = req.Major
	education.Institution = req.Institution
	education.StartDate = req.StartDate
	education.EndDate = req.EndDate

	if err := ec.educationService.UpdateEducation(education); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Education updated successfully"})
}

func (ec educationController) DeleteEducation(c *gin.Context) {
	var req models.EducationDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	education, err := ec.educationService.GetEducationByID(req.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Education not found"})
		return
	}

	if err := ec.educationService.DeleteEducation(education); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Education deleted successfully"})
}

func (ec educationController) GetEducationByID(c *gin.Context) {
	id := c.Param("id")
	uintID, err := utils.StringToUint(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	education, err := ec.educationService.GetEducationByID(uintID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, education)
}

func (ec educationController) GetAllEducations(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	query := models.EducationQuery{
		Degree:      c.Query("degree"),
		Major:       c.Query("major"),
		Institution: c.Query("institution"),
		Page:        utils.StrToInt(c.Query("page"), 1),
		Limit:       utils.StrToInt(c.Query("limit"), 10),
	}

	educations, total, err := ec.educationService.GetAllEducations(userID, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": educations, "total": total})
}
