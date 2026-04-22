package controllers

import (
	"ariesdimasy-portofolio/models"
	"ariesdimasy-portofolio/services"
	"ariesdimasy-portofolio/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExperienceController interface {
	CreateExperience(c *gin.Context)
	UpdateExperience(c *gin.Context)
	DeleteExperience(c *gin.Context)
	GetExperienceByID(c *gin.Context)
	GetAllExperiences(c *gin.Context)
}

type experienceController struct {
	experienceService services.ExperienceService
}

func NewExperienceController(experienceService services.ExperienceService) ExperienceController {
	return &experienceController{experienceService: experienceService}
}

func (ec experienceController) CreateExperience(c *gin.Context) {
	var req models.ExperienceCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Build skills from IDs
	var skills []models.Skill
	for _, skillID := range req.Skill {
		skills = append(skills, models.Skill{ID: skillID})
	}

	experience := models.Experience{
		Title:       req.Title,
		CompanyName: req.CompanyName,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Description: req.Description,
		Skill:       skills,
	}

	if err := ec.experienceService.CreateExperience(&experience); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Experience created successfully"})
}

func (ec experienceController) UpdateExperience(c *gin.Context) {
	id := c.Param("id")
	uintID, err := utils.StringToUint(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req models.ExperienceUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	experience, err := ec.experienceService.GetExperienceByID(uintID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Experience not found"})
		return
	}

	// Build skills from IDs
	var skills []models.Skill
	for _, skillID := range req.Skill {
		skills = append(skills, models.Skill{ID: skillID})
	}

	experience.Title = req.Title
	experience.CompanyName = req.CompanyName
	experience.StartDate = req.StartDate
	experience.EndDate = req.EndDate
	experience.Description = req.Description
	experience.Skill = skills

	if err := ec.experienceService.UpdateExperience(experience); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Experience updated successfully"})
}

func (ec experienceController) DeleteExperience(c *gin.Context) {
	var req models.ExperienceDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	experience, err := ec.experienceService.GetExperienceByID(req.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Experience not found"})
		return
	}

	if err := ec.experienceService.DeleteExperience(experience); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Experience deleted successfully"})
}

func (ec experienceController) GetExperienceByID(c *gin.Context) {
	id := c.Param("id")
	uintID, err := utils.StringToUint(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	experience, err := ec.experienceService.GetExperienceByID(uintID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, experience)
}

func (ec experienceController) GetAllExperiences(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	query := models.ExperienceQuery{
		Title:       c.Query("title"),
		CompanyName: c.Query("company_name"),
		Page:        utils.StrToInt(c.Query("page"), 1),
		Limit:       utils.StrToInt(c.Query("limit"), 10),
	}

	experiences, total, err := ec.experienceService.GetAllExperiences(userID, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": experiences, "total": total})
}
