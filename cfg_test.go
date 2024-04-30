package gocfg_test

import (
	"fmt"
	"github.com/kordar/gocfg"
	"os"
	"testing"
)

func TestDefaultCfg(t *testing.T) {
	gocfg.InitEnv("G", "profile")
	os.Setenv("G_PROFILE", "dev")
	//gocfg.InitConfig("conf/conf.ini")
	gocfg.InitConfigWithDir("default", "conf", "ini", "toml")

	//value := gocfg.GetSystemValue("ee")
	//fmt.Println(value)
	value2 := gocfg.GetGroupSystemValue("default", "ee")
	fmt.Println(value2)

}
