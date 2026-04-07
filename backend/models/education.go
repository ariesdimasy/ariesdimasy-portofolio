package models

import "time"

type Education struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	User        User      `json:"user" gorm:"foreignKey:UserID"`
	Degree      string    `json:"degree" gorm:"not null"`
	Major       string    `json:"major" gorm:"not null"`
	Institution string    `json:"institution" gorm:"not null"`
	StartDate   time.Time `json:"start_date" gorm:"not null"`
	EndDate     time.Time `json:"end_date" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type EducationResponse struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	User        User      `json:"user"`
	Degree      string    `json:"degree"`
	Major       string    `json:"major"`
	Institution string    `json:"institution"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type EducationCreateRequest struct {
	UserID      uint      `json:"user_id" binding:"required"`
	Degree      string    `json:"degree" binding:"required"`
	Major       string    `json:"major" binding:"required"`
	Institution string    `json:"institution" binding:"required"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date" binding:"required"`
}

type EducationUpdateRequest struct {
	ID          uint      `json:"id" binding:"required"`
	Degree      string    `json:"degree" binding:"required"`
	Major       string    `json:"major" binding:"required"`
	Institution string    `json:"institution" binding:"required"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date" binding:"required"`
}

type EducationDeleteRequest struct {
	ID uint `json:"id" binding:"required"`
}
