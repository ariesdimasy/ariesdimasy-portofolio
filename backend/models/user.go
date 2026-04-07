package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Password  string    `json:"password" gorm:"not null"`
	LoginBy   string    `json:"login_by" gorm:"not null, type:enum('email','google')"`
	Avatar    string    `json:"avatar" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	LoginBy   string    `json:"login_by"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserUpdateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserDeleteRequest struct {
	ID uint `json:"id" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Logout struct {
	ID uint `json:"id" binding:"required"`
}

type UserQuery struct {
	Name    string `form:"name"`
	Email   string `form:"email"`
	LoginBy string `form:"login_by"`
	Page    int    `form:"page, default=1"`
	Limit   int    `form:"limit, default=10"`
}
