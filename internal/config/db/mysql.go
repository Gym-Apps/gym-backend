package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "root:mysql123@tcp(127.0.0.1:3306)/dbilanver2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	DB = db
}

// func Migrate(db *gorm.DB) {
// 	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
// 		&model.User{},
// 		&model.UserDetail{},
// 		&model.Adress{},
// 		&model.District{},
// 		&model.Province{},
// 		&model.Category{},
// 		&model.LostPassword{},
// 		&model.ProductDetail{},
// 		&model.ProductState{},
// 		&model.Product{},
// 		&model.Promo{},
// 	)
// }
