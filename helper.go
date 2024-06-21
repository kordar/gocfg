package gocfg

import (
	logger "github.com/kordar/gologger"
	"github.com/spf13/viper"
	"io/ioutil"
	"path"
)

var (
	groups = make(map[string]*viper.Viper)
)

// InitConfigWithSubDir 初始化子目录作为group，适用于多语言场景或不同开发环境
func InitConfigWithSubDir(dir string, ext ...string) {
	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		logger.Fatalf("[gocfg] 读取文件目录失败，pathname=%v, err=%v", dir, err)
		return
	}
	for _, fi := range fis {
		fullname := path.Join(dir, fi.Name())
		if fi.IsDir() {
			InitConfigWithDir(fi.Name(), fullname, ext...)
		}
	}
}

func InitConfig(filepath string) {
	InitDefaultConfig(filepath, "ini")
}

func InitDefaultConfig(filepath string, in string) {
	groups["default"] = loadConfig(viper.New(), []string{filepath}, in)
}

func InitConfigWithDir(group string, parent string, ext ...string) {
	InitCustomerConfigWithDir(viper.New(), group, parent, ext...)
}

func InitCustomerConfigWithDir(v *viper.Viper, group string, parent string, ext ...string) {
	files, err := getAllFile(parent, ext...)
	if err != nil {
		logger.Panic("[gocfg] init cobra fail!")
		return
	}
	groups[group] = loadConfig(v, files, ext...)
}

func loadConfig(v *viper.Viper, files []string, exts ...string) *viper.Viper {
	if len(files) == 0 {
		return v
	}

	mm := make(map[string]bool)
	for _, s := range exts {
		mm[s] = true
	}

	for _, filepath := range files {
		ext := path.Ext(filepath)[1:]
		if !mm[ext] {
			continue
		}
		newViper := viper.New()
		newViper.SetConfigFile(filepath)
		if err := newViper.ReadInConfig(); err == nil {
			_ = v.MergeConfigMap(newViper.AllSettings())
		}
	}

	return v
}

// 递归获取指定目录下的所有文件名
func getAllFile(pathname string, ext ...string) ([]string, error) {
	result := make([]string, 0)

	fis, err := ioutil.ReadDir(pathname)
	if err != nil {
		logger.Errorf("[gocfg] 读取文件目录失败，pathname=%v, err=%v", pathname, err)
		return result, err
	}

	// 所有文件/文件夹
	for _, fi := range fis {
		fullname := path.Join(pathname, fi.Name())
		// 是文件夹则递归进入获取;是文件，则压入数组
		if fi.IsDir() {
			temp, err2 := getAllFile(fullname, ext...)
			if err2 != nil {
				logger.Errorf("[gocfg] 读取文件目录失败,fullname=%v, err=%v", fullname, err)
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
