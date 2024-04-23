package gocfg

import (
	"github.com/spf13/viper"
	"log"
)

var mgr map[string]*viper.Viper

func init() {
	mgr = make(map[string]*viper.Viper)
}

func InitConfig(name string, filepath string, in string) {
	if mgr[name] != nil {
		return
	}
	mgr[name] = viper.New()
	mgr[name].SetConfigFile(filepath)
	mgr[name].SetConfigType(in)
	err := mgr[name].ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// 默认配置
var cfg = viper.New()

func InitDefaultConfig(filepath string, in string) {
	cfg.SetConfigFile(filepath)
	cfg.SetConfigType(in)
	errCfg := cfg.ReadInConfig()
	if errCfg != nil {
		log.Fatal(errCfg)
	}
}

// GetConfig 获取config配置
func GetConfig(name string) *viper.Viper {
	return mgr[name]
}

func GetSystemValue(key string) string {
	return cfg.GetString("system." + key)
}

func GetSettingValue(key string) string {
	return cfg.GetString("setting." + key)
}

func GetSectionValue(section string, key string) string {
	return cfg.GetString(section + "." + key)
}

func GetSectionValueInt(section string, key string) int {
	return cfg.GetInt(section + "." + key)
}

func GetSection(section string) map[string]string {
	return cfg.GetStringMapString(section)
}

func UnmarshalKey(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) error {
	return cfg.UnmarshalKey(key, rawVal, opts...)
}

func GetCfg() *viper.Viper {
	return cfg
}
