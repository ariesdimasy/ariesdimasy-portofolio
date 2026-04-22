package services

import (
	"ariesdimasy-portofolio/models"
	"ariesdimasy-portofolio/repositories"
)

type SosmedService interface {
	CreateSosmed(sosmed *models.Sosmed) error
	UpdateSosmed(sosmed *models.Sosmed) error
	DeleteSosmed(sosmed *models.Sosmed) error
	GetSosmedByID(id uint) (*models.Sosmed, error)
	GetAllSosmeds(userID uint, query models.SosmedQuery) ([]models.Sosmed, error)
}

type sosmedService struct {
	sosmedRepo repositories.SosmedRepository
}

func NewSosmedService(sosmedRepo repositories.SosmedRepository) SosmedService {
	return &sosmedService{sosmedRepo: sosmedRepo}
}

func (ss sosmedService) CreateSosmed(sosmed *models.Sosmed) error {
	return ss.sosmedRepo.CreateSosmed(sosmed)
}

func (ss sosmedService) UpdateSosmed(sosmed *models.Sosmed) error {
	return ss.sosmedRepo.UpdateSosmed(sosmed)
}

func (ss sosmedService) DeleteSosmed(sosmed *models.Sosmed) error {
	return ss.sosmedRepo.DeleteSosmed(sosmed)
}

func (ss sosmedService) GetSosmedByID(id uint) (*models.Sosmed, error) {
	return ss.sosmedRepo.GetSosmedByID(id)
}

func (ss sosmedService) GetAllSosmeds(userID uint, query models.SosmedQuery) ([]models.Sosmed, error) {
	return ss.sosmedRepo.GetAllSosmeds(userID, query)
}
