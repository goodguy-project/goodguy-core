package initialize

import (
	"log"
	"path"

	"github.com/spf13/viper"

	"github.com/goodguy-project/goodguy-core/client/crawl"
	"github.com/goodguy-project/goodguy-core/model"
	"github.com/goodguy-project/goodguy-core/util"
)

func loadViper() {
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	configPath := path.Dir(path.Dir(util.GetFileName()))
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

func MustInit() {
	loadViper()
	model.MustInit()
	createAdmin()
	defaultAdminSet()
	crawl.MustInitClient()
}
