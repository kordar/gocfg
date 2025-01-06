package gocfg

import (
	logger "github.com/kordar/gologger"
	"github.com/spf13/viper"
	"io/ioutil"
	"path"
	"strings"
)

func LoadSubConfig(v *viper.Viper, subDirs []string, exts ...string) *viper.Viper {
	if len(subDirs) == 0 {
		return v
	}

	for _, dir := range subDirs {
		files, err := GetAllFile(dir, exts...)
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

func LoadConfig(v *viper.Viper, files []string, exts ...string) *viper.Viper {
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

	MergeConfig(v, newFiles)
	// -----------
	profile := v.GetString("PROFILE")
	if profile == DEV {
		MergeConfig(v, devFiles)
	}
	if profile == PRO {
		MergeConfig(v, proFiles)
	}
	if profile == TEST {
		MergeConfig(v, testFiles)
	}

	return v
}

func MergeConfig(v *viper.Viper, files []string) {
	for _, filename := range files {
		newViper := viper.New()
		newViper.SetConfigFile(filename)
		if err := newViper.ReadInConfig(); err == nil {
			_ = v.MergeConfigMap(newViper.AllSettings())
		}
	}
}

// 递归获取指定目录下的所有文件名
func GetAllFile(pathname string, ext ...string) ([]string, error) {
	result := make([]string, 0)

	fis, err := ioutil.ReadDir(pathname)
	if err != nil {
		logger.Errorf("[gocfg] failed to read the file directory，pathname=%v, err=%v", pathname, err)
		return result, err
	}

	// 所有文件/文件夹
	for _, fi := range fis {
		fullname := path.Join(pathname, fi.Name())
		// 是文件夹则递归进入获取;是文件，则压入数组
		if fi.IsDir() {
			temp, err2 := GetAllFile(fullname, ext...)
			if err2 != nil {
				logger.Errorf("[gocfg] failed to read the file directory, fullname=%v, err=%v", fullname, err)
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
