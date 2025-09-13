package models

import "gorm.io/gorm"

type Grade struct {
	gorm.Model
	StudentID uint    `json:"student_id"`
	Subject   string  `json:"subject"`
	Score     float64 `json:"score"`
	Student   Student `gorm:"foreignKey:StudentID"`
}