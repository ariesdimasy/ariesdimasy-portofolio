package repositories

import (
	"ariesdimasy-portofolio/models"

	"gorm.io/gorm"
)

type BiodataRepository interface {
	CreateBiodata(biodata *models.Biodata) error
	UpdateBiodata(biodata *models.Biodata) error
	DeleteBiodata(biodata *models.Biodata) error
	GetBiodataByID(id uint) (*models.Biodata, error)
}

type biodataRepository struct {
	db *gorm.DB
}

func NewBiodataRepository(db *gorm.DB) BiodataRepository {
	return &biodataRepository{db: db}
}

func (br biodataRepository) CreateBiodata(biodata *models.Biodata) error {
	return br.db.Create(biodata).Error
}

func (br biodataRepository) UpdateBiodata(biodata *models.Biodata) error {
	return br.db.Save(biodata).Error
}

func (br biodataRepository) DeleteBiodata(biodata *models.Biodata) error {
	return br.db.Delete(biodata).Error
}

func (br biodataRepository) GetBiodataByID(id uint) (*models.Biodata, error) {
	var biodata models.Biodata
	return &biodata, br.db.First(&biodata, id).Error
}
