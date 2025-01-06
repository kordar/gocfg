package gocfg_test

import (
	"bytes"
	"github.com/kordar/gocfg"
	logger "github.com/kordar/gologger"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestInitConfig(t *testing.T) {

	_ = os.Setenv("G_PROFILE", "dev")
	gocfg.InitEnv("G", "profile")
	gocfg.InitConfigWithDir("conf", "ini", "toml", "yaml")

	gocfg.WriteConfigMap(map[string]interface{}{
		"test": map[string]interface{}{
			"aaa": "TTET",
		},
	})

	var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
mm:
 Jacket: leather
 Trousers: denim
age: 35
eyes: brown
beard: true
`)

	gocfg.WriteConfig(yamlExample)

	v := viper.New()
	v.SetConfigType("yaml")
	_ = v.ReadConfig(bytes.NewBuffer(yamlExample))
	gocfg.WriteConfigMap(v.AllSettings())

	sections := gocfg.AllSections()
	logger.Infof("all=%+v", sections)

	testValue := gocfg.Get("test")
	logger.Infof("======testValue=%+v", testValue)
	testSection := gocfg.GetSection("test")
	logger.Infof("======testSection=%+v", testSection)

	mmSection := gocfg.Get("mm")
	logger.Infof("============mmSection=%+v", mmSection)

	type Demo struct {
		M1       string `json:"jacket" mapstructure:"jacket"`
		Trousers string
	}
	d := &Demo{}

	err := gocfg.UnmarshalKey("mm", d)
	logger.Errorf("000000000%+v", err)
	logger.Infof("-----------%v-------%v", d.M1, d.Trousers)
}
