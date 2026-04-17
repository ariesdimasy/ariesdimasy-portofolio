package models

import "time"

type Project struct {
	ID           uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	User         User           `json:"user" gorm:"foreignKey:UserID"`
	UserID       uint           `json:"user_id" gorm:"not null"`
	Title        string         `json:"title" gorm:"not null"`
	Description  string         `json:"description" gorm:"not null"`
	Image        string         `json:"image" gorm:"not null"`
	Link         string         `json:"link" gorm:"not null"`
	Skill        []Skill        `json:"skill" gorm:"one2many:project_skills;foreignKey:ProjectID;references:ID"`
	ProjectImage []ProjectImage `json:"project_image" gorm:"one2many:project_images;foreignKey:ProjectID;references:ID"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
}

type ProjectResponse struct {
	ID           uint           `json:"id"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	Image        string         `json:"image"`
	Link         string         `json:"link"`
	Skill        []Skill        `json:"skill"`
	ProjectImage []ProjectImage `json:"project_image"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type ProjectCreateRequest struct {
	Title        string                      `json:"title" binding:"required"`
	Description  string                      `json:"description" binding:"required"`
	Image        string                      `json:"image" binding:"required"`
	Link         string                      `json:"link" binding:"required"`
	Skill        []uint                      `json:"skill"`
	ProjectImage []ProjectImageCreateRequest `json:"project_image"`
}

type ProjectUpdateRequest struct {
	Title        string                      `json:"title" binding:"required"`
	Description  string                      `json:"description" binding:"required"`
	Image        string                      `json:"image" binding:"required"`
	Link         string                      `json:"link" binding:"required"`
	Skill        []uint                      `json:"skill" `
	ProjectImage []ProjectImageUpdateRequest `json:"project_image"`
}

type ProjectDeleteRequest struct {
	ID uint `json:"id" binding:"required"`
}

type ProjectImage struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	ProjectID uint      `json:"project_id" gorm:"not null"`
	Project   Project   `json:"project" gorm:"foreignKey:ProjectID"`
	Image     string    `json:"image" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type ProjectImageResponse struct {
	ID        uint      `json:"id"`
	ProjectID uint      `json:"project_id"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProjectImageCreateRequest struct {
	ProjectID uint   `json:"project_id" binding:"required"`
	Image     string `json:"image" binding:"required"`
}

type ProjectImageUpdateRequest struct {
	ID    uint   `json:"id" binding:"required"`
	Image string `json:"image" binding:"required"`
}

type ProjectImageDeleteRequest struct {
	ID uint `json:"id" binding:"required"`
}

type ProjectQuery struct {
	Title       string `form:"title"`
	Description string `form:"description"`
	Page        int    `form:"page, default=1"`
	Limit       int    `form:"limit, default=10"`
}
