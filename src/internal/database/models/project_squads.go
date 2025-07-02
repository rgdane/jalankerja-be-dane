package models

import "time"

const TableName = "project_squads"

type ProjectSquad struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ProjectID int64     `gorm:"column:project_id;not null" json:"project_id"`
	SquadID   int64     `gorm:"column:squad_id;not null" json:"squad_id"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`

	HasProject *Project `gorm:"foreignKey:ProjectID;references:ID" json:"has_project,omitempty"`
	HasSquad   *Squad   `gorm:"foreignKey:SquadID;references:ID" json:"has_squad,omitempty"`
}

func (*ProjectSquad) TableName() string {
	return TableName
}
