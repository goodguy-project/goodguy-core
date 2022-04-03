package main

import (
	"log"
	"path"

	"github.com/spf13/viper"

	"github.com/goodguy-project/goodguy-core/cron"
	"github.com/goodguy-project/goodguy-core/model"
	"github.com/goodguy-project/goodguy-core/util"
	"github.com/goodguy-project/goodguy-core/web"
)

func loadViper() {
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	configPath := path.Dir(util.GetFileName())
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func createAdmin() {
	var err error
	member := &model.Member{}
	model.GetDB().Model(&model.Member{}).Where("sid = ?", "admin").First(member)
	if member.Sid == "admin" {
		log.Printf("member admin exists")
		return
	}
	log.Println("member admin creating")
	member = &model.Member{
		Sid:     "admin",
		Name:    "管理员",
		IsAdmin: true,
		Pwd:     util.Hashing(viper.GetString("admin.pwd")),
	}
	err = model.GetDB().Create(member).Error
	if err != nil {
		panic(err)
	}
}

func defaultAdminSet() {
	viper.Set(util.OpenRegisterConfigName, viper.GetBool(util.OpenRegisterConfigName))
	emailConf := util.EmailConfigName
	if viper.GetString(util.EmailConfigName) == "" {
		emailConf = "[]"
	}
	viper.Set(util.EmailConfigName, emailConf)
	err := viper.WriteConfig()
	if err != nil {
		panic(err)
	}
}

func mustInit() {
	loadViper()
	createAdmin()
	defaultAdminSet()
}

func main() {
	mustInit()
	go web.Serve()
	go cron.Serve()
	select {}
}
