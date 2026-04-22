package repositories

import (
	"ariesdimasy-portofolio/models"

	"gorm.io/gorm"
)

type SosmedRepository interface {
	CreateSosmed(sosmed *models.Sosmed) error
	UpdateSosmed(sosmed *models.Sosmed) error
	DeleteSosmed(sosmed *models.Sosmed) error
	GetSosmedByID(id uint) (*models.Sosmed, error)
	GetAllSosmeds(userID uint, query models.SosmedQuery) ([]models.Sosmed, error)
}

type sosmedRepository struct {
	db *gorm.DB
}

func NewSosmedRepository(db *gorm.DB) SosmedRepository {
	return &sosmedRepository{db: db}
}

func (sr sosmedRepository) CreateSosmed(sosmed *models.Sosmed) error {
	return sr.db.Create(sosmed).Error
}

func (sr sosmedRepository) UpdateSosmed(sosmed *models.Sosmed) error {
	return sr.db.Save(sosmed).Error
}

func (sr sosmedRepository) DeleteSosmed(sosmed *models.Sosmed) error {
	return sr.db.Delete(sosmed).Error
}

func (sr sosmedRepository) GetSosmedByID(id uint) (*models.Sosmed, error) {
	var sosmed models.Sosmed
	return &sosmed, sr.db.First(&sosmed, id).Error
}

func (sr sosmedRepository) GetAllSosmeds(userID uint, query models.SosmedQuery) ([]models.Sosmed, error) {
	var sosmeds []models.Sosmed
	dbData := sr.db.Model(&models.Sosmed{}).Where("user_id = ?", userID)

	if query.Name != "" {
		dbData = dbData.Where("name LIKE ?", "%"+query.Name+"%")
	}

	return sosmeds, dbData.Find(&sosmeds).Error
}
