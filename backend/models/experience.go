package models

import "time"

type Experience struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	User        User      `json:"user" gorm:"foreignKey:UserID"`
	Title       string    `json:"title" gorm:"not null"`
	CompanyName string    `json:"company_name" gorm:"not null"`
	StartDate   time.Time `json:"start_date" gorm:"not null"`
	EndDate     time.Time `json:"end_date" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Skill       []Skill   `json:"skill" gorm:"many2many:experience_skills;"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type ExperienceResponse struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	User        User      `json:"user"`
	Title       string    `json:"title"`
	CompanyName string    `json:"company_name"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
	Skill       []Skill   `json:"skill"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ExperienceCreateRequest struct {
	Title       string    `json:"title" binding:"required"`
	CompanyName string    `json:"company_name" binding:"required"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Skill       []uint    `json:"skills" binding:"required"`
}

type ExperienceUpdateRequest struct {
	Title       string    `json:"title" binding:"required"`
	CompanyName string    `json:"company_name" binding:"required"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Skill       []uint    `json:"skills" binding:"required"`
}

type ExperienceDeleteRequest struct {
	ID uint `json:"id" binding:"required"`
}

type ExperienceQuery struct {
	Title       string `form:"title"`
	CompanyName string `form:"company_name"`
	Page        int    `form:"page, default=1"`
	Limit       int    `form:"limit, default=10"`
}
