package services

import (
	"ariesdimasy-portofolio/models"
	"ariesdimasy-portofolio/repositories"
)

type ExperienceService interface {
	CreateExperience(experience *models.Experience) error
	UpdateExperience(experience *models.Experience) error
	DeleteExperience(experience *models.Experience) error
	GetExperienceByID(id uint) (*models.Experience, error)
	GetAllExperiences(userID uint, query models.ExperienceQuery) ([]models.Experience, int64, error)
}

type experienceService struct {
	repo repositories.ExperienceRepository
}

func NewExperienceService(experienceRepository repositories.ExperienceRepository) ExperienceService {
	return &experienceService{repo: experienceRepository}
}

func (es experienceService) CreateExperience(experience *models.Experience) error {
	return es.repo.CreateExperience(experience)
}

func (es experienceService) UpdateExperience(experience *models.Experience) error {
	return es.repo.UpdateExperience(experience)
}

func (es experienceService) DeleteExperience(experience *models.Experience) error {
	return es.repo.DeleteExperience(experience)
}

func (es experienceService) GetExperienceByID(id uint) (*models.Experience, error) {
	return es.repo.GetExperienceByID(id)
}

func (es experienceService) GetAllExperiences(userID uint, query models.ExperienceQuery) ([]models.Experience, int64, error) {
	return es.repo.GetAllExperiences(userID, query)
}
