package services

import (
	"ariesdimasy-portofolio/models"
	"ariesdimasy-portofolio/repositories"
)

type CertificateService interface {
	CreateCertificate(certificate *models.Certificate) error
	UpdateCertificate(certificate *models.Certificate) error
	DeleteCertificate(certificate *models.Certificate) error
	GetCertificateByID(id uint) (*models.Certificate, error)
	GetAllCertificates(userID uint, query models.CertificateQuery) ([]models.Certificate, int64, error)
}

type certificateService struct {
	repo repositories.CertificateRepository
}

func NewCertificateService(certificateRepository repositories.CertificateRepository) CertificateService {
	return &certificateService{repo: certificateRepository}
}

func (cs certificateService) CreateCertificate(certificate *models.Certificate) error {
	return cs.repo.CreateCertificate(certificate)
}

func (cs certificateService) UpdateCertificate(certificate *models.Certificate) error {
	return cs.repo.UpdateCertificate(certificate)
}

func (cs certificateService) DeleteCertificate(certificate *models.Certificate) error {
	return cs.repo.DeleteCertificate(certificate)
}

func (cs certificateService) GetCertificateByID(id uint) (*models.Certificate, error) {
	return cs.repo.GetCertificateByID(id)
}

func (cs certificateService) GetAllCertificates(userID uint, query models.CertificateQuery) ([]models.Certificate, int64, error) {
	return cs.repo.GetAllCertificates(userID, query)
}
