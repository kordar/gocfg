package gocfg

import (
	"errors"
	"github.com/spf13/viper"
)

func GetGroupViper(group string) *viper.Viper {
	return groups[group]
}

func GetGroupSystemValue(group string, key string) string {
	v := groups[group]
	if v == nil {
		return ""
	}
	return v.GetString("system." + key)
}

func GetGroupSettingValue(group string, key string) string {
	v := groups[group]
	if v == nil {
		return ""
	}
	return v.GetString("setting." + key)
}

func GetGroupSectionValue(group string, section string, key string) string {
	v := groups[group]
	if v == nil {
		return ""
	}
	return v.GetString(section + "." + key)
}

func GetGroupSectionValueInt(group string, section string, key string) int {
	v := groups[group]
	if v == nil {
		return 0
	}
	return v.GetInt(section + "." + key)
}

func GetGroupSection(group string, section string) map[string]string {
	v := groups[group]
	if v == nil {
		return map[string]string{}
	}
	return v.GetStringMapString(section)
}

func GroupUnmarshalKey(group string, section string, raw interface{}) error {
	v := groups[group]
	if v == nil {
		return errors.New("not found config")
	}
	return v.UnmarshalKey(section, raw)
}

func GroupGet(group string, key string) interface{} {
	v := groups[group]
	if v == nil {
		return nil
	}
	return v.Get(key)
}
