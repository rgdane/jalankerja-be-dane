package models

import (
	"time"
)

type GeneralDivisi struct {
	ID              int64            `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama            string           `gorm:"size:255;not null" json:"nama"`
	Kode            *string          `gorm:"size:255;uniqueIndex" json:"kode,omitempty"`
	CreatedAt       *time.Time       `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt       time.Time        `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	GeneralProjects []GeneralProject `gorm:"foreignKey:DivisiID;references:Kode" json:"general_projects,omitempty"`
}

func (GeneralDivisi) TableName() string {
	return "general_divisis"
}
