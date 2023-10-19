package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	SgocId           string  `gorm:"default:null" json:"area"`
	Name             string  `gorm:"not null" validate:"required" json:"name"`
	Province         string  `gorm:"default:null" json:"province"`
	District         string  `gorm:"default:null" json:"district"`
	SubDistrict      string  `gorm:"default:null" json:"sub_district"`
	ZipCode          string  `gorm:"default:null" json:"zip_code"`
	Longitude        string  `gorm:"default:null" json:"longitude"`
	Latitude         string  `gorm:"default:null" json:"latitude"`
	Status           string  `gorm:"not null" validate:"required" json:"status"`
	DistanceLocation float32 `gorm:"default:null" json:"distance_location"`
	HistoryAmount    int8    `gorm:"default:null" json:"history_amount"`
	ImageLogo        string  `gorm:"default:null" json:"image_logo"`
}
