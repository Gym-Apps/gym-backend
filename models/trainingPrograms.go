// antreman programları
package models

import "gorm.io/gorm"

type TrainingProgram struct {
	gorm.Model
	Name       string `json:"name"`
	UserID     uint   `json:"user_id"`
	User       User   `json:"user" gorm:"foreignKey:user_id"`
	CoverPhoto string `json:"cover_photo"`
}
