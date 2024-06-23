package gocfg

import (
	"errors"
	"github.com/spf13/viper"
)

func groupName(options ...interface{}) string {
	switch options[0].(type) {
	case string:
		return options[0].(string)
	default:
		return "default"
	}
}

func GetViper(options ...interface{}) *viper.Viper {
	return GetViperObj(groupName(options...))
}

func SetViper(v *viper.Viper, options ...interface{}) {
	name := groupName(options...)
	item := groups[name]
	if item == nil {
		item = NewElement(name)
		item.SetValue(v)
		groups[name] = item
		return
	}
	item.SetValue(v)
}

func WriteConfig(b []byte, options ...interface{}) {
	name := groupName(options...)
	item := groups[name]
	if item != nil {
		item.Write(b)
	}
}

func WriteConfigMap(cfg map[string]interface{}, options ...interface{}) {
	name := groupName(options...)
	item := groups[name]
	if item != nil {
		item.WriteMap(cfg)
	}
}

func UpdateValue(key string, value interface{}, options ...interface{}) {
	name := groupName(options...)
	item := groups[name]
	if item != nil {
		item.Update(key, value)
	}
}

func Get(key string, options ...interface{}) interface{} {
	name := groupName(options...)
	v := GetViperObj(name)
	if v == nil {
		return nil
	}
	return v.Get(key)
}

func GetSystemValue(key string, options ...interface{}) string {
	name := groupName(options...)
	v := GetViperObj(name)
	if v == nil {
		return ""
	}
	return v.GetString("system." + key)
}

func GetSettingValue(key string, options ...interface{}) string {
	name := groupName(options...)
	v := GetViperObj(name)
	if v == nil {
		return ""
	}
	return v.GetString("setting." + key)
}

func GetSectionValue(section string, key string, options ...interface{}) string {
	name := groupName(options...)
	v := GetViperObj(name)
	if v == nil {
		return ""
	}
	return v.GetString(section + "." + key)
}

func GetSectionValueInt(section string, key string, options ...interface{}) int {
	name := groupName(options...)
	v := GetViperObj(name)
	if v == nil {
		return 0
	}
	return v.GetInt(section + "." + key)
}

func GetSection(section string, options ...interface{}) map[string]string {
	name := groupName(options...)
	v := GetViperObj(name)
	if v == nil {
		return map[string]string{}
	}
	return v.GetStringMapString(section)
}

func UnmarshalKey(section string, raw interface{}, options ...interface{}) error {
	name := groupName(options...)
	v := GetViperObj(name)
	if v == nil {
		return errors.New("not found config")
	}
	return v.UnmarshalKey(section, raw)
}

func Sub(key string, options ...interface{}) *viper.Viper {
	name := groupName(options...)
	v := GetViperObj(name)
	if v == nil {
		return nil
	}
	return v.Sub(key)
}

func AllSections(options ...interface{}) map[string]interface{} {
	name := groupName(options...)
	v := GetViperObj(name)
	if v == nil {
		return map[string]interface{}{}
	}
	return v.AllSettings()
}
