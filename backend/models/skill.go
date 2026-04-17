package models

import "time"

type Skill struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"not null;unique"`
	Icon      string    `json:"icon" gorm:"not null"`
	Type      string    `json:"type" gorm:"type:enum('frontend','backend','database')"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type SkillResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SkillCreateRequest struct {
	Name string `json:"name" binding:"required"`
	Icon string `json:"icon" binding:"required"`
	Type string `json:"type" binding:"required,oneof=frontend backend database"`
}

type SkillUpdateRequest struct {
	Name string `json:"name" binding:"required"`
	Icon string `json:"icon" binding:"required"`
	Type string `json:"type" binding:"required"`
}

type SkillDeleteRequest struct {
	ID uint `json:"id" binding:"required"`
}

type SkillQuery struct {
	Name  string `form:"name"`
	Type  string `form:"type"`
	Page  int    `form:"page, default=1"`
	Limit int    `form:"limit, default=10"`
}
