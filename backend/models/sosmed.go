package models

import "time"

type Sosmed struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	Name      string    `json:"name" gorm:"not null"`
	Icon      string    `json:"icon" gorm:"not null"`
	Link      string    `json:"link" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type SosmedResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	User      User      `json:"user"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SosmedCreateRequest struct {
	UserID uint   `json:"user_id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Icon   string `json:"icon" binding:"required"`
	Link   string `json:"link" binding:"required"`
}

type SosmedUpdateRequest struct {
	ID   uint   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Icon string `json:"icon" binding:"required"`
	Link string `json:"link" binding:"required"`
}

type SosmedDeleteRequest struct {
	ID uint `json:"id" binding:"required"`
}
