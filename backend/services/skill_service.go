package services

import (
	"ariesdimasy-portofolio/models"
	"ariesdimasy-portofolio/repositories"
)

type SkillService interface {
	CreateSkill(skill *models.Skill) error
	UpdateSkill(skill *models.Skill) error
	DeleteSkill(skill *models.Skill) error
	GetSkillByID(id uint) (*models.Skill, error)
	GetAllSkills(query models.SkillQuery) ([]models.Skill, error)
}

type skillService struct {
	skillRepo repositories.SkillRepository
}

func NewSkillService(skillRepo repositories.SkillRepository) SkillService {
	return &skillService{skillRepo: skillRepo}
}

func (ss skillService) CreateSkill(skill *models.Skill) error {
	return ss.skillRepo.CreateSkill(skill)
}

func (ss skillService) UpdateSkill(skill *models.Skill) error {
	return ss.skillRepo.UpdateSkill(skill)
}

func (ss skillService) DeleteSkill(skill *models.Skill) error {
	return ss.skillRepo.DeleteSkill(skill)
}

func (ss skillService) GetSkillByID(id uint) (*models.Skill, error) {
	return ss.skillRepo.GetSkillByID(id)
}

func (ss skillService) GetAllSkills(query models.SkillQuery) ([]models.Skill, error) {
	return ss.skillRepo.GetAllSkills(query)
}
