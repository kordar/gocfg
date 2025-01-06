package gocfg

import "github.com/spf13/viper"

type Cfg interface {
	Name() string
	Viper(args ...interface{}) *viper.Viper
	SetViper(v *viper.Viper, args ...interface{})
	Write(b []byte, args ...interface{})
	WriteMap(cfg map[string]interface{}, args ...interface{})
	Update(key string, value interface{}, args ...interface{})
	Get(key string, args ...interface{}) interface{}
	GetSystemValue(key string, args ...interface{}) string
	GetSettingValue(key string, args ...interface{}) string
	GetSectionValue(section string, key string, args ...interface{}) string
	GetSectionValueInt(section string, key string, args ...interface{}) int
	GetSection(section string, args ...interface{}) map[string]string
	UnmarshalKey(section string, raw interface{}, args ...interface{}) error
	Sub(key string, args ...interface{}) *viper.Viper
	AllSettings(args ...interface{}) map[string]interface{}
}
