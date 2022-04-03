package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	const dsn = "root:goodguy@tcp(127.0.0.1:3306)/goodguy_core?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Member{})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db.Debug()
}
