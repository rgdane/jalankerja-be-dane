package models

import (
	"time"

	"gorm.io/datatypes"
)

type Squad struct {
	ID               int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name             string         `gorm:"size:255;not null" json:"name"`
	GeneralProjectID int64          `gorm:"column:general_project_id;not null" json:"general_project_id"`
	Anggota          datatypes.JSON `gorm:"type:json" json:"anggota"`
	CreatedAt        time.Time      `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Squad) TableName() string {
	return "squads"
}
