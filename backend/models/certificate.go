package models

import "time"

type Certificate struct {
	ID             uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID         uint      `json:"user_id" gorm:"not null"`
	User           User      `json:"user" gorm:"foreignKey:UserID"`
	Name           string    `json:"name" gorm:"not null"`
	Organization   string    `json:"organization" gorm:"not null"`
	CredentialID   string    `json:"credential_id" gorm:"not null"`
	CredentialURL  string    `json:"credential_url" gorm:"not null"`
	Image          string    `json:"image"`
	IssueDate      time.Time `json:"issue_date" gorm:"not null"`
	ExpirationDate time.Time `json:"expiration_date" gorm:"not null"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CertificateResponse struct {
	ID             uint      `json:"id"`
	UserID         uint      `json:"user_id"`
	User           User      `json:"user"`
	Name           string    `json:"name"`
	Organization   string    `json:"organization" gorm:"not null"`
	CredentialID   string    `json:"credential_id" gorm:"not null"`
	CredentialURL  string    `json:"credential_url" gorm:"not null"`
	IssueDate      time.Time `json:"issue_date" gorm:"not null"`
	ExpirationDate time.Time `json:"expiration_date" gorm:"not null"`
	Image          string    `json:"image"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CertificateCreateRequest struct {
	UserID         uint      `json:"user_id" binding:"required"`
	Name           string    `json:"name" binding:"required"`
	Organization   string    `json:"organization" gorm:"not null"`
	CredentialID   string    `json:"credential_id" gorm:"not null"`
	CredentialURL  string    `json:"credential_url" gorm:"not null"`
	IssueDate      time.Time `json:"issue_date" gorm:"not null"`
	ExpirationDate time.Time `json:"expiration_date" gorm:"not null"`
	Image          string    `json:"image"`
}

type CertificateUpdateRequest struct {
	ID             uint      `json:"id" binding:"required"`
	Name           string    `json:"name" binding:"required"`
	Organization   string    `json:"organization" gorm:"not null"`
	CredentialID   string    `json:"credential_id" gorm:"not null"`
	CredentialURL  string    `json:"credential_url" gorm:"not null"`
	IssueDate      time.Time `json:"issue_date" gorm:"not null"`
	ExpirationDate time.Time `json:"expiration_date" gorm:"not null"`
	Image          string    `json:"image"`
}

type CertificateDeleteRequest struct {
	ID uint `json:"id" binding:"required"`
}
