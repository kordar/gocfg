package gocfg

import (
	logger "github.com/kordar/gologger"
	"github.com/spf13/viper"
	"io/ioutil"
	"path"
	"strings"
)

var (
	groups = make(map[string]*Element)
)

func GetViperObj(name string) *viper.Viper {
	if groups[name] == nil {
		return nil
	}
	return groups[name].GetValue()
}

func InitConfigWithParentDirG(group string, parent string, ext ...string) {
	fis, err := ioutil.ReadDir(parent)
	if err != nil {
		logger.Fatalf("[gocfg] 读取文件目录失败，pathname=%v, err=%v", parent, err)
		return
	}
	subDirs := make([]string, 0)
	for _, fi := range fis {
		subname := path.Join(parent, fi.Name())
		if fi.IsDir() {
			subDirs = append(subDirs, subname)
		}
	}
	config := loadSubConfig(viper.New(), subDirs, ext...)
	g := NewElement(group)
	g.SetValue(config)
	groups[g.name] = g
}

func InitConfigWithParentDir(parent string, ext ...string) {
	InitConfigWithParentDirG(parent, parent, ext...)
}

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
	config := loadConfig(viper.New(), []string{filepath}, in)
	g := NewElement("default")
	g.SetValue(config)
	groups[g.name] = g
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
	config := loadConfig(v, files, ext...)
	g := NewElement(group)
	g.SetValue(config)
	groups[g.name] = g
}

func loadSubConfig(v *viper.Viper, subDirs []string, exts ...string) *viper.Viper {
	if len(subDirs) == 0 {
		return v
	}

	for _, dir := range subDirs {
		files, err := getAllFile(dir, exts...)
		if err != nil {
			logger.Panic("[gocfg] init cobra fail!")
			continue
		}

		vv := viper.New()
		name := path.Base(dir)
		for _, filename := range files {
			newViper := viper.New()
			newViper.SetConfigFile(filename)
			if err2 := newViper.ReadInConfig(); err2 == nil {
				_ = vv.MergeConfigMap(newViper.AllSettings())
			}
		}

		v.Set(name, vv.AllSettings())
	}

	return v
}

func loadConfig(v *viper.Viper, files []string, exts ...string) *viper.Viper {
	if len(files) == 0 {
		return v
	}

	mm := make(map[string]bool)
	for _, s := range exts {
		mm[s] = true
	}

	newFiles := make([]string, 0)
	devFiles := make([]string, 0)
	proFiles := make([]string, 0)
	testFiles := make([]string, 0)

	for _, filename := range files {
		ext := path.Ext(filename)[1:]
		if !mm[ext] {
			continue
		}
		if strings.Contains(filename, "-dev.") {
			devFiles = append(devFiles, filename)
			continue
		}
		if strings.Contains(filename, "-pro.") {
			proFiles = append(proFiles, filename)
			continue
		}
		if strings.Contains(filename, "-test.") {
			testFiles = append(testFiles, filename)
			continue
		}
		newFiles = append(newFiles, filename)
	}

	mergeConfig(v, newFiles)
	// -----------
	profile := viper.GetString("PROFILE")
	if profile == DEV {
		mergeConfig(v, devFiles)
	}
	if profile == PRO {
		mergeConfig(v, proFiles)
	}
	if profile == TEST {
		mergeConfig(v, testFiles)
	}

	return v
}

func mergeConfig(v *viper.Viper, files []string) {
	for _, filename := range files {
		newViper := viper.New()
		newViper.SetConfigFile(filename)
		if err := newViper.ReadInConfig(); err == nil {
			_ = v.MergeConfigMap(newViper.AllSettings())
		}
	}
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
