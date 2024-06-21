package gocfg_test

import (
	"fmt"
	"github.com/kordar/gocfg"
	logger "github.com/kordar/gologger"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestDefaultCfg(t *testing.T) {
	gocfg.InitEnv("G", "profile")
	_ = os.Setenv("G_PROFILE", "")
	//gocfg.InitConfig("conf/conf.ini")
	gocfg.InitConfigWithDir("default", "conf", "ini", "toml", "yaml")
	value := gocfg.GetGroupSection("default", "system")
	logger.Infof("===============%v", value)

	//value := gocfg.GetSystemValue("ee")
	//fmt.Println(value)
	//value2 := gocfg.GetSection("a")
	//fmt.Println(value2)
	//
	type Demo struct {
		Hai int    `json:"hai"`
		Mi  string `json:"mi"`
	}
	d := &Demo{}
	//section := gocfg.GetSection("mm")
	//fmt.Printf("======%+v", section)
	err := gocfg.UnmarshalKey("mm", d)
	fmt.Printf("error=%v", err)
	fmt.Printf("--------------%v\n", d)

}

func TestT22(t *testing.T) {
	v := viper.New()
	v.SetConfigName("./conf/tt.toml")
	v.SetConfigType("toml")
	v.AddConfigPath(".")
	err := v.ReadInConfig() // 查找并读取配置文件
	logger.Errorf("----------error = %v", err)
	str := v.GetStringMap("d")

	logger.Infof("==========%+v", str)
}
