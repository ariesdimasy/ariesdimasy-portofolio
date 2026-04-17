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
	GetAllCertificates(userID uint, query models.CertificateQuery) ([]models.Certificate, error)
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

func (cr certificateRepository) GetAllCertificates(userID uint, query models.CertificateQuery) ([]models.Certificate, error) {
	var certificates []models.Certificate
	var total int64

	errCount := cr.db.Model(&models.Certificate{}).Where("user_id = ?", userID).Count(&total)

	if errCount.Error != nil {
		return nil, errCount.Error
	}

	if total == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	offset := (query.Page - 1) * query.Limit
	dbData := cr.db.Model(&models.Certificate{}).Where("user_id = ?", userID).Offset(offset).Limit(query.Limit)

	if query.Name != "" {
		dbData = dbData.Where("name LIKE ?", "%"+query.Name+"%")
	}

	if err := dbData.Find(&certificates).Error; err != nil {
		return nil, err
	}

	return certificates, nil
}
