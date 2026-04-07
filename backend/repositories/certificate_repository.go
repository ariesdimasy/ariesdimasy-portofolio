package repositories

import (
	"ariesdimasy-portofolio/models"

	"gorm.io/gorm"
)

type CertificateRepository interface {
	CreateCertificate(certificate *models.Certificate) error
	UpdateCertificate(certificate *models.Certificate) error
	DeleteCertificate(certificate *models.Certificate) error
	GetCertificateByID(id uint) (*models.Certificate, error)
	GetAllCertificates() ([]models.Certificate, error)
}

type certificateRepository struct {
	db *gorm.DB
}

func NewCertificateRepository(db *gorm.DB) CertificateRepository {
	return &certificateRepository{db: db}
}

func (cr certificateRepository) CreateCertificate(certificate *models.Certificate) error {
	return cr.db.Create(certificate).Error
}

func (cr certificateRepository) UpdateCertificate(certificate *models.Certificate) error {
	return cr.db.Save(certificate).Error
}

func (cr certificateRepository) DeleteCertificate(certificate *models.Certificate) error {
	return cr.db.Delete(certificate).Error
}

func (cr certificateRepository) GetCertificateByID(id uint) (*models.Certificate, error) {
	var certificate models.Certificate
	return &certificate, cr.db.First(&certificate, id).Error
}

func (cr certificateRepository) GetAllCertificates() ([]models.Certificate, error) {
	var certificates []models.Certificate
	return certificates, cr.db.Find(&certificates).Error
}
