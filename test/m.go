package main

import (
	"github.com/kordar/gocfg"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	//gocfg.InitConfig("test", "conf/conf.ini", "ini")
	//
	//f := gocfg.ProfileFlag{}
	//gocfg.BindFlagValue("profile", f)
	//fmt.Printf("---------%+v\n", viper.GetString("profile"))
	//flag.Parse()

	//viper.SetEnvPrefix("aa")

	gocfg.InitEnv("AA", "ID", "id", "name", "c1", "c2")
	_ = os.Setenv("id", "4444")
	_ = os.Setenv("c2", "5555")
	//_ = os.Setenv("c1", "45657")
	log.Println("---------------", viper.Get("id"), viper.Get("c1"))
}
