package repositories

import (
	"ariesdimasy-portofolio/models"

	"gorm.io/gorm"
)

type educationRepository struct {
	db *gorm.DB
}

func NewEducationRepository(db *gorm.DB) EducationRepository {
	return &educationRepository{db: db}
}

type EducationRepository interface {
	CreateEducation(education *models.Education) error
	UpdateEducation(education *models.Education) error
	DeleteEducation(education *models.Education) error
	GetEducationByID(id uint) (*models.Education, error)
	GetAllEducations(userID uint, query models.EducationQuery) ([]models.Education, int64, error)
}

func (er educationRepository) CreateEducation(education *models.Education) error {
	return er.db.Create(education).Error
}

func (er educationRepository) UpdateEducation(education *models.Education) error {
	return er.db.Save(education).Error
}

func (er educationRepository) DeleteEducation(education *models.Education) error {
	return er.db.Delete(education).Error
}

func (er educationRepository) GetEducationByID(id uint) (*models.Education, error) {
	var education models.Education
	return &education, er.db.First(&education, id).Error
}

func (er educationRepository) GetAllEducations(userID uint, query models.EducationQuery) ([]models.Education, int64, error) {
	var educations []models.Education
	var total int64

	errCount := er.db.Model(&models.Education{}).Where("user_id = ?", userID).Count(&total)
	if errCount.Error != nil {
		return nil, 0, errCount.Error
	}

	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}

	dbData := er.db.Model(&models.Education{}).Where("user_id = ?", userID)

	if query.Degree != "" {
		dbData = dbData.Where("degree LIKE ?", "%"+query.Degree+"%")
	}

	if query.Major != "" {
		dbData = dbData.Where("major LIKE ?", "%"+query.Major+"%")
	}

	if query.Institution != "" {
		dbData = dbData.Where("institution LIKE ?", "%"+query.Institution+"%")
	}

	if err := dbData.Where("user_id = ?", userID).Find(&educations).Error; err != nil {
		return nil, total, err
	}

	return educations, total, nil
}
