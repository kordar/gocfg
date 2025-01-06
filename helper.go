package gocfg

import (
	"errors"
	logger "github.com/kordar/gologger"
	"github.com/spf13/viper"
	"io/ioutil"
	"path"
)

func GetViperObj() *viper.Viper {
	return GetSnippet("default").GetValue()
}

func InitConfigWithParentDir(parent string, ext ...string) {
	fis, err := ioutil.ReadDir(parent)
	if err != nil {
		logger.Fatalf("[gocfg] failed to read the file directory，pathname=%v, err=%v", parent, err)
		return
	}
	subDirs := make([]string, 0)
	for _, fi := range fis {
		subname := path.Join(parent, fi.Name())
		if fi.IsDir() {
			subDirs = append(subDirs, subname)
		}
	}
	viperObj := GetViperObj()
	config := LoadSubConfig(viperObj, subDirs, ext...)
	GetSnippet("default").SetValue(config)
}

func InitConfig(filepath string) {
	InitDefaultConfig(filepath, "ini")
}

func InitDefaultConfig(filepath string, in string) {
	viperObj := GetViperObj()
	config := LoadConfig(viperObj, []string{filepath}, in)
	GetSnippet("default").SetValue(config)
}

func InitConfigWithDir(parent string, ext ...string) {
	viperObj := GetViperObj()
	InitCustomerConfigWithDir(viperObj, parent, ext...)
}

func InitCustomerConfigWithDir(v *viper.Viper, parent string, ext ...string) {
	files, err := GetAllFile(parent, ext...)
	if err != nil {
		logger.Panic("[gocfg] init cobra fail!")
		return
	}
	config := LoadConfig(v, files, ext...)
	GetSnippet("default").SetValue(config)
}

func SetViper(v *viper.Viper) {
	GetSnippet("default").SetValue(v)
}

func WriteConfig(b []byte) {
	GetSnippet("default").Write(b)
}

func WriteConfigMap(cfg map[string]interface{}) {
	GetSnippet("default").WriteMap(cfg)
}

func UpdateValue(key string, value interface{}) {
	GetSnippet("default").Update(key, value)
}

func Get(key string) interface{} {
	v := GetViperObj()
	if v == nil {
		return nil
	}
	return v.Get(key)
}

func GetSystemValue(key string) string {
	v := GetViperObj()
	if v == nil {
		return ""
	}
	return v.GetString("system." + key)
}

func GetSettingValue(key string) string {
	v := GetViperObj()
	if v == nil {
		return ""
	}
	return v.GetString("setting." + key)
}

func GetSectionValue(section string, key string) string {
	v := GetViperObj()
	if v == nil {
		return ""
	}
	return v.GetString(section + "." + key)
}

func GetSectionValueInt(section string, key string) int {
	v := GetViperObj()
	if v == nil {
		return 0
	}
	return v.GetInt(section + "." + key)
}

func GetSection(section string) map[string]string {
	v := GetViperObj()
	if v == nil {
		return map[string]string{}
	}
	return v.GetStringMapString(section)
}

func UnmarshalKey(section string, raw interface{}) error {
	v := GetViperObj()
	if v == nil {
		return errors.New("not found config")
	}
	return v.UnmarshalKey(section, raw)
}

func Sub(key string) *viper.Viper {
	v := GetViperObj()
	if v == nil {
		return nil
	}
	return v.Sub(key)
}

func AllSections() map[string]interface{} {
	v := GetViperObj()
	if v == nil {
		return map[string]interface{}{}
	}
	return v.AllSettings()
}
