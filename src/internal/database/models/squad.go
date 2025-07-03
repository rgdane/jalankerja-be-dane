package models

import (
	"time"
)

type Squad struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"size:255;not null" json:"name"`
	Description string    `gorm:"size:255" json:"description"`
	Captain     *int64    `gorm:"index" json:"captain"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	HasCaptain   *User          `gorm:"foreignKey:Captain;references:ID" json:"has_captain,omitempty"`
	ProjectSquad []ProjectSquad `gorm:"foreignKey:SquadID;references:ID" json:"has_project"`
}

func (Squad) TableName() string {
	return "squads"
}
