// öğun tablosu.
package models

import "gorm.io/gorm"

type Meal struct {
	gorm.Model
	NutritionalValue []MealNutrition `json:"meal_nutrientation"`
	UserID           uint            `json:"user_id"`
	User             User            `json:"user" gorm:"foreignKey:user_id"`
	Name             string
}

func (m *Meal) AfterFind(db *gorm.DB) error {
	return nil
}
