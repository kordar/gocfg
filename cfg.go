package gocfg

import (
	log "github.com/kordar/gologger"
	"github.com/spf13/viper"
	"io/ioutil"
	"path"
)

func InitConfig(filepath string) {
	InitDefaultConfig(filepath, "ini")
}

func InitDefaultConfig(filepath string, in string) {
	groupItem := GroupItem{
		Group:      "default",
		Files:      []string{filepath},
		ExtType:    in,
		DriverName: "viper",
	}
	AddCobraByGroup(groupItem)
}

func GetSystemValue(key string) string {
	cobra := GetCobra("default", "ini", "viper")
	return cobra.GetCfg().GetString("system." + key)
}

func GetSettingValue(key string) string {
	cobra := GetCobra("default", "ini", "viper")
	return cobra.GetCfg().GetString("setting." + key)
}

func GetSectionValue(section string, key string) string {
	cobra := GetCobra("default", "ini", "viper")
	return cobra.GetCfg().GetString(section + "." + key)
}

func GetSectionValueInt(section string, key string) int {
	cobra := GetCobra("default", "ini", "viper")
	return cobra.GetCfg().GetInt(section + "." + key)
}

func GetSection(section string) map[string]string {
	cobra := GetCobra("default", "ini", "viper")
	return cobra.GetCfg().GetStringMapString(section)
}

func UnmarshalKey(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) error {
	cfg := GetCfg()
	return cfg.UnmarshalKey(key, rawVal, opts...)
}

func GetCfg() *viper.Viper {
	cobra := GetCobra("default", "ini", "viper")
	return cobra.GetCfg().GetInstance().(*viper.Viper)
}

// GetAllFile 递归获取指定目录下的所有文件名
func GetAllFile(pathname string, ext ...string) ([]string, error) {
	result := make([]string, 0)

	fis, err := ioutil.ReadDir(pathname)
	if err != nil {
		log.Errorf("[gocfg] 读取文件目录失败，pathname=%v, err=%v", pathname, err)
		return result, err
	}

	// 所有文件/文件夹
	for _, fi := range fis {
		fullname := path.Join(pathname, fi.Name())
		// 是文件夹则递归进入获取;是文件，则压入数组
		if fi.IsDir() {
			temp, err2 := GetAllFile(fullname, ext...)
			if err2 != nil {
				log.Errorf("[gocfg] 读取文件目录失败,fullname=%v, err=%v", fullname, err)
				return result, err2
			}
			result = append(result, temp...)
		} else {
			suffix := path.Ext(fullname)[1:]
			flag := false
			for _, s := range ext {
				if suffix == s {
					flag = true
					break
				}
			}
			if flag {
				result = append(result, fullname)
			}
		}
	}

	return result, nil
}

func InitConfigWithDir(group string, parent string, ext ...string) {
	files, err := GetAllFile(parent, ext...)
	//log.Infof("----------------%+v", files)
	if err != nil {
		log.Panic("[gocfg] init cobra fail!")
		return
	}
	configItems := make([]ConfigItem, 0)
	for _, file := range files {
		configItem := ConfigItem{
			Group:      group,
			Filename:   file,
			ExtType:    path.Ext(file)[1:],
			DriverName: "viper",
		}
		configItems = append(configItems, configItem)
	}
	groupItemList := ToGroupItemList(configItems)
	for _, groupItem := range groupItemList {
		AddCobraByGroup(groupItem)
	}
}

// InitConfigWithSubDir 初始化子目录作为group，适用于多语言场景或不同开发环境
func InitConfigWithSubDir(dir string, ext ...string) {
	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalf("[gocfg] 读取文件目录失败，pathname=%v, err=%v", dir, err)
		return
	}
	for _, fi := range fis {
		fullname := path.Join(dir, fi.Name())
		if fi.IsDir() {
			InitConfigWithDir(fi.Name(), fullname, ext...)
		}
	}
}

func GetGroupSystemValue(group string, key string) string {
	cobra := GetCobraWithKey(group, "system."+key)
	if cobra == nil {
		return ""
	}
	return cobra.GetCfg().GetString("system." + key)
}

func GetGroupSettingValue(group string, key string) string {
	cobra := GetCobraWithKey(group, "setting."+key)
	if cobra == nil {
		return ""
	}
	return cobra.GetCfg().GetString("setting." + key)
}

func GetGroupSectionValue(group string, section string, key string) string {
	cobra := GetCobraWithKey(group, section+"."+key)
	if cobra == nil {
		return ""
	}
	return cobra.GetCfg().GetString(section + "." + key)
}

func GetGroupSectionValueInt(group string, section string, key string) int {
	cobra := GetCobraWithKey(group, section+"."+key)
	if cobra == nil {
		return 0
	}
	return cobra.GetCfg().GetInt(section + "." + key)
}

func GetGroupSection(group string, section string) map[string]string {
	cobra := GetCobraWithKey(group, section)
	if cobra == nil {
		return map[string]string{}
	}
	return cobra.GetCfg().GetStringMapString(section)
}
