package services

import (
	"ariesdimasy-portofolio/models"
	"ariesdimasy-portofolio/repositories"
)

type EducationService interface {
	CreateEducation(education *models.Education) error
	UpdateEducation(education *models.Education) error
	DeleteEducation(education *models.Education) error
	GetEducationByID(id uint) (*models.Education, error)
	GetAllEducations(userID uint, query models.EducationQuery) ([]models.Education, int64, error)
}

type educationService struct {
	educationRepo repositories.EducationRepository
}

func NewEducationService(educationRepo repositories.EducationRepository) EducationService {
	return &educationService{educationRepo: educationRepo}
}

func (es educationService) CreateEducation(education *models.Education) error {
	return es.educationRepo.CreateEducation(education)
}

func (es educationService) UpdateEducation(education *models.Education) error {
	return es.educationRepo.UpdateEducation(education)
}

func (es educationService) DeleteEducation(education *models.Education) error {
	return es.educationRepo.DeleteEducation(education)
}

func (es educationService) GetEducationByID(id uint) (*models.Education, error) {
	return es.educationRepo.GetEducationByID(id)
}

func (es educationService) GetAllEducations(userID uint, query models.EducationQuery) ([]models.Education, int64, error) {
	return es.educationRepo.GetAllEducations(userID, query)
}
