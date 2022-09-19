package db

import (
	"context"
	"log"
	"time"

	"github.com/Gym-Apps/gym-backend/models"
	_ "github.com/newrelic/go-agent/v3/integrations/nrmysql"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Time time.Duration

func configSetup() {
	viper.SetConfigFile(`../config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("gymapp started with debug mode")
	}

	Time = time.Duration(viper.GetInt("context.timeout")) * time.Second
}

func Connect(app *newrelic.Application) *gorm.DB {

	configSetup()
	dbHost := viper.GetString(`database.host`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	dsn := dbUser + ":" + dbPass + dbHost + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "nrmysql",
		DSN:        dsn,
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	DB = db

	gormTransactionTrace := app.StartTransaction("GORM Operation")
	gormTransactionContext := newrelic.NewContext(context.Background(), gormTransactionTrace)
	db = db.WithContext(gormTransactionContext)

	db.Debug().Find(&models.User{})

	gormTransactionTrace.End()

	if viper.GetBool(`debug`) {
		Migrate()
	}

	return db
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
