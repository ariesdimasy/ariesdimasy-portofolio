package controllers

import (
	"ariesdimasy-portofolio/models"
	"ariesdimasy-portofolio/services"
	"ariesdimasy-portofolio/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProjectController interface {
	CreateProject(c *gin.Context)
	UpdateProject(c *gin.Context)
	DeleteProject(c *gin.Context)
	GetProjectByID(c *gin.Context)
	GetAllProjects(c *gin.Context)
}

type projectController struct {
	projectService services.ProjectService
}

func NewProjectController(projectService services.ProjectService) ProjectController {
	return &projectController{projectService: projectService}
}

func (pc projectController) CreateProject(c *gin.Context) {
	var req models.ProjectCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Build skills from IDs
	var skills []models.Skill
	for _, skillID := range req.Skill {
		skills = append(skills, models.Skill{ID: skillID})
	}

	// Build project images
	var projectImages []models.ProjectImage
	for _, img := range req.ProjectImage {
		projectImages = append(projectImages, models.ProjectImage{
			Image: img.Image,
		})
	}

	project := models.Project{
		Title:        req.Title,
		Description:  req.Description,
		Image:        req.Image,
		Link:         req.Link,
		Skill:        skills,
		ProjectImage: projectImages,
	}

	if err := pc.projectService.CreateProject(&project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project created successfully"})
}

func (pc projectController) UpdateProject(c *gin.Context) {
	id := c.Param("id")
	uintID, err := utils.StringToUint(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req models.ProjectUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project, err := pc.projectService.GetProjectByID(uintID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	// Build skills from IDs
	var skills []models.Skill
	for _, skillID := range req.Skill {
		skills = append(skills, models.Skill{ID: skillID})
	}

	// Build project images from update requests
	var projectImages []models.ProjectImage
	for _, img := range req.ProjectImage {
		projectImages = append(projectImages, models.ProjectImage{
			ID:    img.ID,
			Image: img.Image,
		})
	}

	project.Title = req.Title
	project.Description = req.Description
	project.Image = req.Image
	project.Link = req.Link
	project.Skill = skills
	project.ProjectImage = projectImages

	if err := pc.projectService.UpdateProject(project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project updated successfully"})
}

func (pc projectController) DeleteProject(c *gin.Context) {
	var req models.ProjectDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project, err := pc.projectService.GetProjectByID(req.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	if err := pc.projectService.DeleteProject(project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}

func (pc projectController) GetProjectByID(c *gin.Context) {
	id := c.Param("id")
	uintID, err := utils.StringToUint(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project, err := pc.projectService.GetProjectByID(uintID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}

func (pc projectController) GetAllProjects(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	query := models.ProjectQuery{
		Title:       c.Query("title"),
		Description: c.Query("description"),
		Page:        utils.StrToInt(c.Query("page"), 1),
		Limit:       utils.StrToInt(c.Query("limit"), 10),
	}

	projects, total, err := pc.projectService.GetAllProjects(userID, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": projects, "total": total})
}
