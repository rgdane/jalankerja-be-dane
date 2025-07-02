package models

import (
	"time"
)

const TableNameProject = "projects"

// Project mapped from table <projects>
type Project struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Nama      string    `gorm:"column:nama;not null" json:"nama"`
	Kode      string    `gorm:"column:kode;not null" json:"kode"`
	Deskripsi string    `gorm:"column:deskripsi" json:"deskripsi"`
	StartDate time.Time `gorm:"column:start_date" json:"start_date"`
	EndDate   time.Time `gorm:"column:end_date" json:"end_date"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`

	ProjectSquad []ProjectSquad `gorm:"foreignKey:ProjectID;references:ID" json:"project_squad"`
}

// TableName Project's table name
func (*Project) TableName() string {
	return TableNameProject
}
