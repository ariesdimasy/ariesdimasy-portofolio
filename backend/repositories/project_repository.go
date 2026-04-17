package repositories

import (
	"ariesdimasy-portofolio/models"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	CreateProject(project *models.Project) error
	UpdateProject(project *models.Project) error
	DeleteProject(project *models.Project) error
	GetProjectByID(id uint) (*models.Project, error)
	GetAllProjects(userID uint, query models.ProjectQuery) ([]models.Project, int64, error)
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{db: db}
}

func (pr projectRepository) CreateProject(project *models.Project) error {
	return pr.db.Create(project).Error
}

func (pr projectRepository) UpdateProject(project *models.Project) error {
	return pr.db.Save(project).Error
}

func (pr projectRepository) DeleteProject(project *models.Project) error {
	return pr.db.Delete(project).Error
}

func (pr projectRepository) GetProjectByID(id uint) (*models.Project, error) {
	var project models.Project
	return &project, pr.db.Preload("Skill").Preload("ProjectImage").First(&project, id).Error
}

func (pr projectRepository) GetAllProjects(userID uint, query models.ProjectQuery) ([]models.Project, int64, error) {
	var projects []models.Project
	var total int64

	errCount := pr.db.Model(&models.Project{}).Where("user_id = ?", userID).Count(&total)
	if errCount.Error != nil {
		return nil, 0, errCount.Error
	}

	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}

	offset := (query.Page - 1) * query.Limit
	dbData := pr.db.Model(&models.Project{}).Where("user_id = ?", userID).Offset(offset).Limit(query.Limit)

	if query.Title != "" {
		dbData = dbData.Where("title LIKE ?", "%"+query.Title+"%")
	}

	if query.Description != "" {
		dbData = dbData.Where("description LIKE ?", "%"+query.Description+"%")
	}

	if err := dbData.Preload("Skill").Preload("ProjectImage").Find(&projects).Error; err != nil {
		return nil, total, err
	}

	return projects, total, nil
}
