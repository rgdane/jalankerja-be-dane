package models

import (
	"time"

	"gorm.io/datatypes"
)

type User struct {
	ID              int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name            string         `gorm:"size:255" json:"name"`
	Email           string         `gorm:"size:255;uniqueIndex" json:"email"`
	EmailVerifiedAt *time.Time     `gorm:"column:email_verified_at" json:"email_verified_at"`
	Password        string         `gorm:"size:255" json:"-"`
	RememberToken   *string        `gorm:"size:100" json:"remember_token,omitempty"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	CustomFields    datatypes.JSON `json:"custom_fields"`
	AvatarURL       *string        `gorm:"column:avatar_url;size:255" json:"avatar_url,omitempty"`
}

func (User) TableName() string {
	return "users"
}
