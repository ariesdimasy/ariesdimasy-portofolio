package services

import (
	"ariesdimasy-portofolio/models"
	"ariesdimasy-portofolio/repositories"
)

type BiodataService interface {
	CreateBiodata(biodata *models.Biodata) error
	UpdateBiodata(biodata *models.Biodata) error
	DeleteBiodata(biodata *models.Biodata) error
	GetBiodataByID(id uint) (*models.Biodata, error)
}

type biodataService struct {
	repo repositories.BiodataRepository
}

func NewBiodataService(biodataRepository repositories.BiodataRepository) BiodataService {
	return &biodataService{repo: biodataRepository}
}

func (bs biodataService) CreateBiodata(biodata *models.Biodata) error {
	return bs.repo.CreateBiodata(biodata)
}

func (bs biodataService) UpdateBiodata(biodata *models.Biodata) error {
	return bs.repo.UpdateBiodata(biodata)
}

func (bs biodataService) DeleteBiodata(biodata *models.Biodata) error {
	return bs.repo.DeleteBiodata(biodata)
}

func (bs biodataService) GetBiodataByID(id uint) (*models.Biodata, error) {
	return bs.repo.GetBiodataByID(id)
}
