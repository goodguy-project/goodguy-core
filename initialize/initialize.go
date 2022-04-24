package initialize

import (
	"github.com/goodguy-project/goodguy-core/core/constant"
	"log"

	"github.com/goodguy-project/goodguy-core/model"
	"github.com/goodguy-project/goodguy-core/util"
	"github.com/goodguy-project/goodguy-core/util/conf"
)

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
		Pwd:     util.Hashing(conf.Viper().GetString("admin.pwd")),
	}
	err = model.GetDB().Create(member).Error
	if err != nil {
		panic(err)
	}
}

func defaultAdminSet() {
	conf.Viper().Set(constant.OpenRegisterConfigName, conf.Viper().GetBool(constant.OpenRegisterConfigName))
	emailConf := constant.EmailConfigName
	if conf.Viper().GetString(constant.EmailConfigName) == "" {
		emailConf = "[]"
	}
	conf.Viper().Set(constant.EmailConfigName, emailConf)
	err := conf.Viper().WriteConfig()
	if err != nil {
		panic(err)
	}
}

func init() {
	defaultAdminSet()
}

func MustInit() {
	model.MustInit()
	createAdmin()
}
