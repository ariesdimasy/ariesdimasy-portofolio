package services

import (
	"ariesdimasy-portofolio/models"
	"ariesdimasy-portofolio/repositories"
)

type ProjectService interface {
	CreateProject(project *models.Project) error
	UpdateProject(project *models.Project) error
	DeleteProject(project *models.Project) error
	GetProjectByID(id uint) (*models.Project, error)
	GetAllProjects(userID uint, query models.ProjectQuery) ([]models.Project, int64, error)
}

type projectService struct {
	projectRepo repositories.ProjectRepository
}

func NewProjectService(projectRepo repositories.ProjectRepository) ProjectService {
	return &projectService{projectRepo: projectRepo}
}

func (ps projectService) CreateProject(project *models.Project) error {
	return ps.projectRepo.CreateProject(project)
}

func (ps projectService) UpdateProject(project *models.Project) error {
	return ps.projectRepo.UpdateProject(project)
}

func (ps projectService) DeleteProject(project *models.Project) error {
	return ps.projectRepo.DeleteProject(project)
}

func (ps projectService) GetProjectByID(id uint) (*models.Project, error) {
	return ps.projectRepo.GetProjectByID(id)
}

func (ps projectService) GetAllProjects(userID uint, query models.ProjectQuery) ([]models.Project, int64, error) {
	return ps.projectRepo.GetAllProjects(userID, query)
}
