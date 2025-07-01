package models

import (
	"time"
)

type Squad struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Squad) TableName() string {
	return "squads"
}
