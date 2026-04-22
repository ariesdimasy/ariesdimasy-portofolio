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
	GetAllSkills(query models.SkillQuery) ([]models.Skill, error)
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

func (sr skillRepository) GetAllSkills(query models.SkillQuery) ([]models.Skill, error) {
	var skills []models.Skill

	dbData := sr.db.Model(&models.Skill{})

	if query.Name != "" {
		dbData = dbData.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if query.Type != "" {
		dbData = dbData.Where("type = ?", query.Type)
	}

	if err := dbData.Find(&skills).Error; err != nil {
		return nil, err
	}

	return skills, nil
}
