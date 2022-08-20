package db

import (
	"github.com/Gym-Apps/gym-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var debug = true
var DB *gorm.DB

func Init() {
	dsn := "root:mysql123@tcp(127.0.0.1:3306)/gymapp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	DB = db

	if debug {
		Migrate()
	}

}

func Migrate() {
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.User{},
		&models.BodyMeasurement{},
		&models.Movement{},
		&models.TrainingProgram{},
		&models.Training_movent{},
		&models.NutritionalValue{},
		&models.Meal{},
		&models.MealNutrition{},
	)
}
