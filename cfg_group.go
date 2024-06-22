package gocfg

import (
	"errors"
	"github.com/spf13/viper"
)

func GetGroupViper(groupName string) *viper.Viper {
	return GetViperObj(groupName)
}

func SetGroupViper(groupName string, v *viper.Viper) {
	item := groups[groupName]
	if item == nil {
		item = NewElement(groupName)
		item.SetValue(v)
		groups[groupName] = item
		return
	}
	item.SetValue(v)
}

func WriteGroupConfig(groupName string, b []byte) {
	item := groups[groupName]
	if item != nil {
		item.Write(b)
	}
}

func WriteGroupConfigMap(groupName string, cfg map[string]interface{}) {
	item := groups[groupName]
	if item != nil {
		item.WriteMap(cfg)
	}
}

func UpdateGroupValue(groupName string, key string, value interface{}) {
	item := groups[groupName]
	if item != nil {
		item.Update(key, value)
	}
}

func GetGroupSystemValue(groupName string, key string) string {
	v := GetViperObj(groupName)
	if v == nil {
		return ""
	}
	return v.GetString("system." + key)
}

func GetGroupSettingValue(groupName string, key string) string {
	v := GetViperObj(groupName)
	if v == nil {
		return ""
	}
	return v.GetString("setting." + key)
}

func GetGroupSectionValue(groupName string, section string, key string) string {
	v := GetViperObj(groupName)
	if v == nil {
		return ""
	}
	return v.GetString(section + "." + key)
}

func GetGroupSectionValueInt(groupName string, section string, key string) int {
	v := GetViperObj(groupName)
	if v == nil {
		return 0
	}
	return v.GetInt(section + "." + key)
}

func GetGroupSection(groupName string, section string) map[string]string {
	v := GetViperObj(groupName)
	if v == nil {
		return map[string]string{}
	}
	return v.GetStringMapString(section)
}

func GroupUnmarshalKey(groupName string, section string, raw interface{}) error {
	v := GetViperObj(groupName)
	if v == nil {
		return errors.New("not found config")
	}
	return v.UnmarshalKey(section, raw)
}

func GroupGet(groupName string, key string) interface{} {
	v := GetViperObj(groupName)
	if v == nil {
		return nil
	}
	return v.Get(key)
}

func GroupSub(groupName string, key string) *viper.Viper {
	v := GetViperObj(groupName)
	if v == nil {
		return nil
	}
	return v.Sub(key)
}
