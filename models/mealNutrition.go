// öğünler ile besinlerin birleştiği tablo
package models

type MealNutrition struct {
	MealID             uint             `json:"meal_id"`
	Meal               Meal             `json:"meal" gorm:"foreignKey:meal_id"`
	NutritionalValueID uint             `json:"nutritional_value_id"`
	NutritionalValue   NutritionalValue `json:"nutritional_value" gorm:"foreignKey:nutritional_value_id"`
	Gram               int32            `json:"gram"`
}
