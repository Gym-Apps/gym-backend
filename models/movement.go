// spor hareketleri tablosu
package models

import "gorm.io/gorm"

type MovementType int8

const (
	Chest MovementType = iota + 1
	Biceps
	Triceps
	Back     // sırt
	Shoulder // omuz
	Leg      // bacak
	Abdomen  // karın
	Heating  // ısınma
)

// reps and sets
type Movement struct {
	gorm.Model
	Name             string       `json:"name"`
	MovementType     MovementType `json:"movement_type"`
	MovementTypeName string       `json:"movement_type_name" gorm:"-"`
}

func (m *Movement) AfterFind(db *gorm.DB) error {
	switch m.MovementType {
	case Chest:
		m.MovementTypeName = "göğüs"
	case Biceps:
		m.MovementTypeName = "Ön kol"
	case Triceps:
		m.MovementTypeName = "Arka kol"
	case Back:
		m.MovementTypeName = "Sırt"
	case Shoulder:
		m.MovementTypeName = "Omuz"
	case Leg:
		m.MovementTypeName = "Bacak"
	case Abdomen:
		m.MovementTypeName = "Karın"
	case Heating:
		m.MovementTypeName = "Isınma"
	}
	return nil
}
