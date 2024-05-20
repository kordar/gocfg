# gocfg 

> 通过集成多种配置三方库实现配置操作统计界面，目前仅实现viper集成



- 初始化配置文件

```go 
func InitConfig(filepath string)  // 初始化ini配置文件 
func InitDefaultConfig(filepath string, in string) // 初始化自定义后缀配置文件 ini toml yaml 等
func InitConfigWithDir(group string, parent string, ext ...string) // 初始化某个目录下的所有ext扩展的文件，group目录名称，parent目录地址，ext扩展名
// InitConfigWithSubDir 初始化子目录作为group，适用于多语言场景或不同开发环境
func InitConfigWithSubDir(dir string, ext ...string)
```

- 基本操作

```go

func GetSystemValue(key string) string 

func GetSettingValue(key string) string 

func GetSectionValue(section string, key string) 

func GetSectionValueInt(section string, key string) int 

func GetSection(section string) map[string]string 

func UnmarshalKey(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) error

func GetCfg() *viper.Viper 




func GetGroupSystemValue(group string, key string) string {
	cobra := GetCobraWithKey(group, "system."+key)
	if cobra == nil {
		return ""
	}
	return cobra.GetCfg().GetString("system." + key)
}

func GetGroupSettingValue(group string, key string) string {
	cobra := GetCobraWithKey(group, "setting."+key)
	if cobra == nil {
		return ""
	}
	return cobra.GetCfg().GetString("setting." + key)
}

func GetGroupSectionValue(group string, section string, key string) string {
	cobra := GetCobraWithKey(group, section+"."+key)
	if cobra == nil {
		return ""
	}
	return cobra.GetCfg().GetString(section + "." + key)
}

func GetGroupSectionValueInt(group string, section string, key string) int {
	cobra := GetCobraWithKey(group, section+"."+key)
	if cobra == nil {
		return 0
	}
	return cobra.GetCfg().GetInt(section + "." + key)
}

func GetGroupSection(group string, section string) map[string]string {
	cobra := GetCobraWithKey(group, section)
	if cobra == nil {
		return map[string]string{}
	}
	return cobra.GetCfg().GetStringMapString(section)
}
```


