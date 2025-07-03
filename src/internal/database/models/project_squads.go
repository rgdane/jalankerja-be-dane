package models

import "time"

const TableName = "project_squads"

type ProjectSquad struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	ProjectID int64     `gorm:"not null" json:"project_id"`
	SquadID   int64     `gorm:"not null" json:"squad_id"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`

	HasProject *Project `gorm:"foreignKey:ProjectID;references:ID;constraint:OnDelete:CASCADE" json:"has_project,omitempty"`
	HasSquad   *Squad   `gorm:"foreignKey:SquadID;references:ID;constraint:OnDelete:CASCADE" json:"has_squad,omitempty"`
}

func (*ProjectSquad) TableName() string {
	return TableName
}
