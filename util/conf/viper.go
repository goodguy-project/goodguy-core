package conf

import (
	"path"
	"runtime"
	"sync"

	"github.com/spf13/viper"
)

var (
	once sync.Once
	vp   = viper.New()
)

func getFileName() string {
	_, fileName, _, _ := runtime.Caller(1)
	return fileName
}

func Viper() *viper.Viper {
	once.Do(func() {
		vp.SetConfigName("config.yaml")
		vp.SetConfigType("yaml")
		configPath := path.Dir(path.Dir(getFileName()))
		vp.AddConfigPath(configPath)
		err := vp.ReadInConfig()
		if err != nil {
			panic(err)
		}
	})
	return vp
}
