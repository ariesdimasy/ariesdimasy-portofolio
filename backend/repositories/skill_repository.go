package repositories

import (
	"ariesdimasy-portofolio/models"

	"gorm.io/gorm"
)

type SkillRepository interface {
	CreateSkill(skill *models.Skill) error
	UpdateSkill(skill *models.Skill) error
	DeleteSkill(skill *models.Skill) error
	GetSkillByID(id uint) (*models.Skill, error)
	GetAllSkills() ([]models.Skill, error)
}

type skillRepository struct {
	db *gorm.DB
}

func NewSkillRepository(db *gorm.DB) SkillRepository {
	return &skillRepository{db: db}
}

func (sr skillRepository) CreateSkill(skill *models.Skill) error {
	return sr.db.Create(skill).Error
}

func (sr skillRepository) UpdateSkill(skill *models.Skill) error {
	return sr.db.Save(skill).Error
}

func (sr skillRepository) DeleteSkill(skill *models.Skill) error {
	return sr.db.Delete(skill).Error
}

func (sr skillRepository) GetSkillByID(id uint) (*models.Skill, error) {
	var skill models.Skill
	return &skill, sr.db.First(&skill, id).Error
}

func (sr skillRepository) GetAllSkills() ([]models.Skill, error) {
	var skills []models.Skill
	return skills, sr.db.Find(&skills).Error
}
