package models

import (
	"time"
)

type GeneralProject struct {
	ID             int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama           string    `gorm:"size:255" json:"nama"`
	Deskripsi      string    `gorm:"type:text" json:"deskripsi"`
	TanggalMulai   time.Time `gorm:"type:date" json:"tanggal_mulai"`
	TanggalSelesai time.Time `gorm:"type:date" json:"tanggal_selesai"`
	UserID         int64     `json:"user_id"`
	DivisiID       string    `gorm:"size:255" json:"divisi_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	HasSquad   []Squad              `gorm:"foreignKey:GeneralProjectID" json:"squad"`
	HasMembers []GeneralProjectUser `gorm:"foreignKey:GeneralProjectID" json:"members"`
	HasDivisi  *GeneralDivisi       `gorm:"foreignKey:DivisiID;references:Kode" json:"divisi"`
}

func (GeneralProject) TableName() string {
	return "general_projects"

}
