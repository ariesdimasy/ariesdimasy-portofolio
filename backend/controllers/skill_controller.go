package controllers

import (
	"ariesdimasy-portofolio/models"
	"ariesdimasy-portofolio/services"
	"ariesdimasy-portofolio/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SkillController interface {
	CreateSkill(c *gin.Context)
	UpdateSkill(c *gin.Context)
	DeleteSkill(c *gin.Context)
	GetSkillByID(c *gin.Context)
	GetAllSkills(c *gin.Context)
}

type skillController struct {
	skillService services.SkillService
}

func NewSkillController(skillService services.SkillService) SkillController {
	return &skillController{skillService: skillService}
}

func (sc skillController) CreateSkill(c *gin.Context) {
	var req models.SkillCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	skill := models.Skill{
		Name: req.Name,
		Icon: req.Icon,
		Type: req.Type,
	}

	if err := sc.skillService.CreateSkill(&skill); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Skill created successfully"})
}

func (sc skillController) UpdateSkill(c *gin.Context) {
	id := c.Param("id")
	uintID, err := utils.StringToUint(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req models.SkillUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	skill, err := sc.skillService.GetSkillByID(uintID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Skill not found"})
		return
	}

	skill.Name = req.Name
	skill.Icon = req.Icon
	skill.Type = req.Type

	if err := sc.skillService.UpdateSkill(skill); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Skill updated successfully"})
}

func (sc skillController) DeleteSkill(c *gin.Context) {
	var req models.SkillDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	skill, err := sc.skillService.GetSkillByID(req.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Skill not found"})
		return
	}

	if err := sc.skillService.DeleteSkill(skill); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Skill deleted successfully"})
}

func (sc skillController) GetSkillByID(c *gin.Context) {
	id := c.Param("id")
	uintID, err := utils.StringToUint(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	skill, err := sc.skillService.GetSkillByID(uintID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, skill)
}

func (sc skillController) GetAllSkills(c *gin.Context) {
	query := models.SkillQuery{
		Name: c.Query("name"),
		Type: c.Query("type"),
	}

	skills, err := sc.skillService.GetAllSkills(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": skills})
}
