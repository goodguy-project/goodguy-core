package model

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = nil
)

func MustInit() {
	var err error
	const dsn = "root:goodguy@tcp(127.0.0.1:3306)/goodguy_core?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Member{}, &MemberContestRecord{}, &SubscribeLog{})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	if db == nil {
		panic(errors.New("init model first"))
	}
	return db.Debug()
}
