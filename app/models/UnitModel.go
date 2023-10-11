package models

import (
	"gorm.io/gorm"
)

type Unit struct {
	gorm.Model
	ProjectId string `gorm:"not null" validate:"required" json:"project_id"`
	Name      string `gorm:"not null" validate:"required" json:"name"`
	Area      string `gorm:"default:null" json:"area"`
	Zone      string `gorm:"default:null" json:"zone"`
	Soi       string `gorm:"default:null" json:"soi"`
	Longitude string `gorm:"default:null" json:"longitude"`
	Latitude  string `gorm:"default:null" json:"latitude"`
	Remark    string `gorm:"default:null" json:"remark"`
	Status    string `gorm:"not null" validate:"required" json:"status"`
}
