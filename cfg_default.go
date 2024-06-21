package gocfg

import "github.com/spf13/viper"

func GetSystemValue(key string) string {
	return GetGroupSystemValue("default", key)
}

func GetSettingValue(key string) string {
	return GetGroupSettingValue("default", key)
}

func GetSectionValue(section string, key string) string {
	return GetGroupSectionValue("default", section, key)
}

func GetSectionValueInt(section string, key string) int {
	return GetGroupSectionValueInt("default", section, key)
}

func GetSection(section string) map[string]string {
	return GetGroupSection("default", section)
}

func UnmarshalKey(key string, rawVal interface{}) error {
	return GroupUnmarshalKey("default", key, rawVal)
}

func GetViper() *viper.Viper {
	return GetGroupViper("default")
}

func Get(key string) interface{} {
	return GroupGet("default", key)
}
