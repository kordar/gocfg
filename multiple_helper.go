package gocfg

import (
	logger "github.com/kordar/gologger"
	"github.com/spf13/viper"
	"io/ioutil"
	"path"
)

var handles = map[string]*Snippet{}

// InitConfigWithSubDir 初始化子目录作为group，适用于多语言场景或不同开发环境
func InitConfigWithSubDir(dir string, ext ...string) {
	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		logger.Fatalf("[gocfg] failed to read the file directory，pathname=%v, err=%v", dir, err)
		return
	}
	for _, fi := range fis {
		fullname := path.Join(dir, fi.Name())
		if fi.IsDir() {
			// InitConfigWithDir(fullname, ext...)
			files, err := GetAllFile(fullname, ext...)
			name := fi.Name()
			if err != nil {
				logger.Panic("[gocfg] init snippet fail!")
				return
			}
			snippet := GetSnippet(name)
			config := LoadConfig(snippet.GetValue(), files, ext...)
			snippet.SetValue(config)
		}
	}
}

func GetSnippet(name string) *Snippet {
	if handles[name] == nil {
		handles[name] = NewSnippet(name, nil)
	}
	return handles[name]
}

func SetViperG(v *viper.Viper, name string) {
	snippet := GetSnippet(name)
	snippet.SetValue(v)
}

func WriteConfigG(b []byte, name string) {
	snippet := GetSnippet(name)
	snippet.Write(b)
}

func WriteConfigMapG(cfg map[string]interface{}, name string) {
	snippet := GetSnippet(name)
	snippet.WriteMap(cfg)
}

func UpdateValueG(key string, value interface{}, name string) {
	snippet := GetSnippet(name)
	snippet.Update(key, value)
}

func GetG(key string, name string) interface{} {
	snippet := GetSnippet(name)
	return snippet.GetValue().Get(key)
}

func GetSystemValueG(key string, name string) string {
	snippet := GetSnippet(name)
	return snippet.GetValue().GetString("system." + key)
}

func GetSettingValueG(key string, name string) string {
	snippet := GetSnippet(name)
	return snippet.GetValue().GetString("setting." + key)
}

func GetSectionValueG(section string, key string, name string) string {
	snippet := GetSnippet(name)
	return snippet.GetValue().GetString(section + "." + key)
}

func GetSectionValueIntG(section string, key string, name string) int {
	snippet := GetSnippet(name)
	return snippet.GetValue().GetInt(section + "." + key)
}

func GetSectionG(section string, name string) map[string]string {
	snippet := GetSnippet(name)
	return snippet.GetValue().GetStringMapString(section)
}

func UnmarshalKeyG(section string, raw interface{}, name string) error {
	snippet := GetSnippet(name)
	return snippet.GetValue().UnmarshalKey(section, raw)
}

func SubG(key string, name string) *viper.Viper {
	snippet := GetSnippet(name)
	return snippet.GetValue().Sub(key)
}

func AllSectionsG(name string) map[string]interface{} {
	snippet := GetSnippet(name)
	return snippet.GetValue().AllSettings()
}
