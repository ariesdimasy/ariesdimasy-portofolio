package repositories

import (
	"ariesdimasy-portofolio/models"

	"gorm.io/gorm"
)

type ExperienceRepository interface {
	CreateExperience(experience *models.Experience) error
	UpdateExperience(experience *models.Experience) error
	DeleteExperience(experience *models.Experience) error
	GetExperienceByID(id uint) (*models.Experience, error)
	GetAllExperiences(userID uint, query models.ExperienceQuery) ([]models.Experience, error)
}

type experienceRepository struct {
	db *gorm.DB
}

func NewExperienceRepository(db *gorm.DB) ExperienceRepository {
	return &experienceRepository{db: db}
}

func (er experienceRepository) CreateExperience(experience *models.Experience) error {
	return er.db.Create(experience).Error
}

func (er experienceRepository) UpdateExperience(experience *models.Experience) error {
	return er.db.Save(experience).Error
}

func (er experienceRepository) DeleteExperience(experience *models.Experience) error {
	return er.db.Delete(experience).Error
}

func (er experienceRepository) GetExperienceByID(id uint) (*models.Experience, error) {
	var experience models.Experience
	return &experience, er.db.Preload("Skill").First(&experience, id).Error
}

func (er experienceRepository) GetAllExperiences(userID uint, query models.ExperienceQuery) ([]models.Experience, error) {
	var experiences []models.Experience
	var total int64
	errCount := er.db.Model(&models.Experience{}).Where("user_id = ?", userID).Count(&total)
	if errCount.Error != nil {
		return nil, errCount.Error
	}

	offset := (query.Page - 1) * query.Limit
	dbData := er.db.Model(&models.Experience{}).Where("user_id = ?", userID).Offset(offset).Limit(query.Limit)

	if query.Title != "" {
		dbData = dbData.Where("title LIKE ?", "%"+query.Title+"%")
	}

	if query.CompanyName != "" {
		dbData = dbData.Where("company_name LIKE ?", "%"+query.CompanyName+"%")
	}

	if err := dbData.Preload("Skill").Find(&experiences).Error; err != nil {
		return nil, err
	}

	return experiences, nil
}
