package model

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB = nil
	once sync.Once
)

func MustInit() {
	once.Do(func() {
		var err error
		const dsn = "root:goodguy@tcp(goodguy-mysql:3306)/goodguy_core?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			var err2 error
			const hystrix = "root:goodguy@tcp(127.0.0.1:9854)/goodguy_core?charset=utf8mb4&parseTime=True&loc=Local"
			db, err2 = gorm.Open(mysql.Open(hystrix), &gorm.Config{})
			if err2 != nil {
				fmt.Printf("err2: %v", err)
				panic(err)
			}
		}
		err = db.AutoMigrate(&Member{}, &MemberContestRecord{}, &SubscribeLog{})
		if err != nil {
			panic(err)
		}
	})
}

func GetDB() *gorm.DB {
	if db == nil {
		MustInit()
	}
	return db.Debug()
}
