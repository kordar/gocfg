package gocfg

import (
	"bytes"
	"io"
	"strings"
	"time"
)

type GoCfg interface {
	WriteConfig()                 // 将当前的 Viper 配置写入预定义的路径（如果存在）。如果没有预定义的路径，则报错。如果配置文件已存在，则会覆盖当前的配置文件。
	SafeWriteConfig()             // 将当前的 Viper 配置写入预定义的路径。如果没有预定义的路径，则报错。如果配置文件已存在，则不会覆盖当前的配置文件。
	WriteConfigAs()               // 将当前的 Viper 配置写入指定的文件路径。如果给定的文件已存在，则会覆盖该文件。
	SafeWriteConfigAs()           // 将当前的 Viper 配置写入指定的文件路径。如果给定的文件已存在，则不会覆盖该文件。 根据经验，标记为 safe 的所有操作都不会覆盖任何文件，只会在文件不存在时创建，而默认行为是创建或截断（truncate）文件。
	WatchConfig()                 // 可选地，您还可以为 Viper 提供一个在每次更改发生时运行的函数
	ReadConfig(bytes.Buffer)      // 从 io.Reader 读取配置
	Set(string, interface{})      // 设置覆盖值
	RegisterAlias(string, string) // 注册和使用别名

	// AutomaticEnv os.Setenv("SPF_ID", "13") // 通常在应用程序外部进行
	AutomaticEnv()                                 // 您可以告诉Viper在读取环境变量时使用前缀
	BindEnv(...string) error                       // 接受一个或多个参数。第一个参数是键名，其余参数是要绑定到此键的环境变量的名称。如果提供了多个参数，它们将按指定的顺序优先。环境变量的名称区分大小写。如果未提供ENV变量名，则Viper将自动假设ENV变量与以下格式匹配：前缀 + “_” +键名（全部大写）。当您显式提供ENV变量名（第二个参数）时，它不会自动添加前缀。例如，如果第二个参数是"id"，Viper将查找ENV变量 “ID”。
	SetEnvPrefix(string)                           // AutomaticEnv是一个强大的辅助功能，特别是与SetEnvPrefix结合使用时。当（AutomaticEnv）被调用时，每次进行viper.Get请求时，Viper都会检查是否存在相应的环境变量。它将应用以下规则：如果设置了EnvPrefix，它将检查是否存在一个与键名相匹配的以大写形式和前缀（如果设置）作为前缀的环境变量名称。
	SetEnvKeyReplacer(...string) *strings.Replacer // 允许您使用strings.Replacer对象来在一定程度上重写Env键。这在您希望在Get()调用中使用-或其他字符，但希望您的环境变量使用_作为分隔符时非常有用。
	AllowEmptyEnv(bool)                            // 默认情况下，空环境变量被视为未设置，并将回退到下一个配置源。要将空环境变量视为已设置，请使用AllowEmptyEnv方法。

	// Get 基本操作
	Get(key string) interface{}
	GetBool(key string) bool
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetIntSlice(key string) []int
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringSlice(key string) []string
	GetStringMapStringSlice(key string) map[string][]string
	GetTime(key string) time.Time
	GetUint(key string) uint
	GetUint16(key string) uint16
	GetUint32(key string) uint32
	GetUint64(key string) uint64
	InConfig(key string) bool // InConfig函数检查给定的键（或别名）是否在配置文件中。
	GetDuration(key string) time.Duration
	IsSet(key string) bool // IsSet 检查键是否在任何数据位置被设置。对于键而言，IsSet函数不区分大小写。
	AllSettings() map[string]interface{}

	MergeConfig(in io.Reader) error // MergeConfig函数将新配置与现有配置合并。
	MergeConfigMap(cfg map[string]interface{}) error
	MergeInConfig() error // MergeInConfig 函数将新的配置与现有配置合并。

	Sub(key string) interface{}

	Unmarshal(rawVal interface{}) error
	UnmarshalKey(key string, rawVal interface{}) error

	KeyDelimiter(v string)
}

func A() {

}

//
//var driver ViperDriver = NewDefaultViper()
//
//func SetDriver(v ViperDriver) {
//	driver = v
//}
//
//func GetViper() *viper.Viper {
//	return driver.GetViper()
//}
//
//func InitConfigWithDir(group string, parent string, ext ...string) {
//	files, err := GetAllFile(parent, ext...)
//	if err != nil {
//		logger.Panic("[gocfg] init cobra fail!")
//		return
//	}
//	configItems := make([]ConfigItem, 0)
//	for _, file := range files {
//		configItem := ConfigItem{
//			Group:      group,
//			Filename:   file,
//			ExtType:    path.Ext(file)[1:],
//			DriverName: "viper",
//		}
//		configItems = append(configItems, configItem)
//	}
//	groupItemList := ToGroupItemList(configItems)
//	for _, groupItem := range groupItemList {
//		AddCobraByGroup(groupItem)
//	}
//}
//
//// InitConfigWithSubDir 初始化子目录作为group，适用于多语言场景或不同开发环境
//func InitConfigWithSubDir(dir string, ext ...string) {
//	fis, err := ioutil.ReadDir(dir)
//	if err != nil {
//		logger.Fatalf("[gocfg] 读取文件目录失败，pathname=%v, err=%v", dir, err)
//		return
//	}
//	for _, fi := range fis {
//		fullname := path.Join(dir, fi.Name())
//		if fi.IsDir() {
//			InitConfigWithDir(fi.Name(), fullname, ext...)
//		}
//	}
//}
//
//func InitConfig(filepath string) {
//	InitDefaultConfig(filepath, "ini")
//}
//
//func InitDefaultConfig(filepath string, in string) {
//	groupItem := GroupItem{
//		Group:      "default",
//		Files:      []string{filepath},
//		ExtType:    in,
//		DriverName: "viper",
//	}
//	AddCobraByGroup(groupItem)
//}

//
//func GetGroupSystemValue(group string, key string) string {
//	cobra := GetCobraWithKey(group, "system."+key)
//	if cobra == nil {
//		return ""
//	}
//	return cobra.GetCfg().GetString("system." + key)
//}
//
//func GetGroupSettingValue(group string, key string) string {
//	cobra := GetCobraWithKey(group, "setting."+key)
//	if cobra == nil {
//		return ""
//	}
//	return cobra.GetCfg().GetString("setting." + key)
//}
//
//func GetGroupSectionValue(group string, section string, key string) string {
//	cobra := GetCobraWithKey(group, section+"."+key)
//	if cobra == nil {
//		return ""
//	}
//	return cobra.GetCfg().GetString(section + "." + key)
//}
//
//func GetGroupSectionValueInt(group string, section string, key string) int {
//	cobra := GetCobraWithKey(group, section+"."+key)
//	if cobra == nil {
//		return 0
//	}
//	return cobra.GetCfg().GetInt(section + "." + key)
//}
//
//func GetGroupSection(group string, section string) map[string]string {
//	cobra := GetCobraWithKey(group, section)
//	if cobra == nil {
//		return map[string]string{}
//	}
//	return cobra.GetCfg().GetStringMapString(section)
//}
//
//func GroupUnmarshalKey(group string, section string, raw interface{}) error {
//	cobra := GetCobraWithKey(group, section)
//	if cobra == nil {
//		return errors.New("not found config")
//	}
//	return cobra.GetCfg().UnmarshalKey(section, raw)
//}
//
//func GetGroupDriver(group string) Driver {
//	return GetCobra(group, "ini", "viper").GetCfg()
//}
//
//func GroupGet(group string, key string) interface{} {
//	cobra := GetCobraWithKey(group, key)
//	if cobra == nil {
//		return nil
//	}
//	return cobra.GetCfg().Get(key)
//}
