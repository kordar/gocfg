package gocfg_test

import (
	"bytes"
	"fmt"
	"github.com/kordar/gocfg"
	logger "github.com/kordar/gologger"
	"github.com/spf13/viper"
	"os"
	"testing"
	"time"
)

func TestDefaultCfg(t *testing.T) {
	gocfg.InitEnv("G", "profile")
	_ = os.Setenv("G_PROFILE", "dev")
	//gocfg.InitConfig("conf/conf.ini")
	gocfg.InitConfigWithDir("default", "conf", "ini", "toml", "yaml")
	value := gocfg.GetSection("system", nil)
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

func TestRW(t *testing.T) {
	gocfg.InitConfigWithDir("default", "conf", "ini", "toml", "yaml")
	var name = "default"

	var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
mm:
 jacket: leather
 trousers: denim
age: 35
eyes: brown
beard: true
`)

	v := viper.New()
	v.SetConfigType("yaml")
	err := v.ReadConfig(bytes.NewBuffer(yamlExample))
	if err != nil {
		logger.Error(err)
	}
	get := v.Get("mm")
	logger.Infof("------------------%v", get)

	//viper.ReadConfig(bytes.NewBuffer(yamlExample))

	for i := 0; i < 50; i++ {
		read(name, "mm")
		//if i%5 == 0 {
		//	write(name, map[string]interface{}{"mm": map[string]interface{}{"cc": fmt.Sprintf("dd-%d", i), "hai": int64(1234)}})
		//}
		if i%21 == 0 {
			gocfg.UpdateValue(name, "mm", map[string]interface{}{"dddd": "we3"})
		}
		//read(name, "mm")
	}
	time.Sleep(30 * time.Second)
}

func read(name string, key string) {
	section := gocfg.GetSection(name, key)
	logger.Infof("read======%v", section)
}

func write(name string, m map[string]interface{}) {
	gocfg.WriteConfigMap(m, name)
}

func TestInitConfigWithParentDir(t *testing.T) {
	gocfg.InitConfigWithParentDir("language", "ini")
	logger.Infof("===================%v", gocfg.AllSections("language"))
	logger.Infof("===================%v", gocfg.GetSectionValue("zh_CN", "dict.this is tom!!", "language"))
	v := viper.New()
	v.Set("en", map[string]interface{}{"aa": "EEEEE"})
	gocfg.SetViper(v, "language")
	logger.Infof("===================%v", gocfg.AllSections("language"))
	logger.Infof("===================%v", gocfg.GetSectionValue("en", "aa", "language"))

}

func TestT003(t *testing.T) {
	var a = func(b ...interface{}) {
		switch b[0].(type) {
		case string:
			logger.Info("string =", b)
			break
		case nil:
			logger.Infof("xxxxeeeeeeee")
			break
		default:
			logger.Infof("ccccc")
		}
	}
	a(nil)
}
