package models

import (
	"time"
)

type GeneralProjectUser struct {
	ID               int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	GeneralProjectID int64      `gorm:"column:general_project_id;not null" json:"general_project_id"`
	UserID           int64      `gorm:"column:user_id;not null" json:"user_id"`
	CreatedAt        *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt        time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	Project GeneralProject `gorm:"foreignKey:GeneralProjectID;constraint:OnDelete:CASCADE,OnUpdate:NO ACTION" json:"project"`
	User    User           `gorm:"foreignKey:UserID" json:"user"`
}

func (GeneralProjectUser) TableName() string {
	return "general_project_user"
}
