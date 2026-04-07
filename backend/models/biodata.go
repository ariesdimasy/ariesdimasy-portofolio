package models

import "time"

type Biodata struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	Headline  string    `json:"headline" gorm:"not null"`
	About     string    `json:"about" gorm:"not null"`
	Address   string    `json:"address" gorm:"not null"`
	Phone     string    `json:"phone" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type BiodataResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	User      User      `json:"user"`
	Headline  string    `json:"headline"`
	About     string    `json:"about"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BiodataCreateRequest struct {
	UserID   uint   `json:"user_id" binding:"required"`
	Headline string `json:"headline" binding:"required"`
	About    string `json:"about" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type BiodataUpdateRequest struct {
	ID       uint   `json:"id" binding:"required"`
	Headline string `json:"headline" binding:"required"`
	About    string `json:"about" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type BiodataDeleteRequest struct {
	ID uint `json:"id" binding:"required"`
}
