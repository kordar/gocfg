package gocfg

//import (
//	"github.com/spf13/viper"
//	"time"
//)
//
//type ViperDriver interface {
//	Name() string
//	GetViper() *viper.Viper
//}
//
//type DefaultDriver struct {
//	v *viper.Viper
//}
//
//func NewDefaultViper() *DefaultDriver {
//	return &DefaultDriver{v: viper.GetViper()}
//}
//
//func (d DefaultDriver) Name() string {
//	return "default"
//}
//
//func (d DefaultDriver) GetViper() *viper.Viper {
//	return d.v
//}
//
//type Driver interface {
//	GetName() string
//	LoadFiles(groupItem GroupItem)
//	GetInstance() interface{}
//	Get(key string) interface{}
//	GetString(key string) string
//	GetBool(key string) bool
//	GetInt(key string) int
//	GetInt32(key string) int32
//	GetInt64(key string) int64
//	GetUint(key string) uint
//	GetUint32(key string) uint32
//	GetUint64(key string) uint64
//	GetFloat64(key string) float64
//	GetTime(key string) time.Time
//	GetDuration(key string) time.Duration
//	GetIntSlice(key string) []int
//	GetStringSlice(key string) []string
//	GetStringMap(key string) map[string]interface{}
//	GetStringMapString(key string) map[string]string
//	GetStringMapStringSlice(key string) map[string][]string
//	GetSizeInBytes(key string) uint
//	IsSet(key string) bool
//	AllSettings() map[string]interface{}
//	Unmarshal(rawVal interface{}) error
//	UnmarshalExact(rawVal interface{}) error
//	UnmarshalKey(key string, rawVal interface{}) error
//}
//
//func CreateDriver(name string) Driver {
//	return &ViperDriver{}
//}
