# gocfg

`gocfg`通过对`viper`进行二次封装，可进行多配置文件合并，配置分组等功能，提供独立对象进行配置初始化。

## 配置文件初始化

- 扫描指定目录下的所有文件

```golang
func InitConfigWithParentDir(parent string, ext ...string)  // 仅扫描父目录下的子目录及子目录下所有文件
func InitConfigWithDir(parent string, ext ...string) // 当前文件下的所有文件及目录
func InitCustomerConfigWithDir(v *viper.Viper, parent string, ext ...string) // 自定viper对象
```

- 指定配置文件

```go
func InitConfig(filepath string)  // 默认支持ini文件
func InitDefaultConfig(filepath string, in string) // 自定义文件类型
```

- 通过子目录进行分组

```go
func InitConfigWithSubDir(dir string, ext ...string)  // 将目录下的一级子目录作为分组key，生成多个配置文件
```

一般适用于多语言场景或不同开发环境

默认支持的扩展文件`ini,yml,yaml,toml`，可自定义扩展。

## 独立配置对象

实现`Cfg`接口即为独立配置文件，默认提供`LocalCfg`和`MultipleCfg`

```go
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
```



## 环境变量设置

```go
_ = os.Setenv("G_PROFILE", "dev") 
gocfg.InitEnv("G", "profile")  // 绑定g_profile作为环境变量
```

设置文件名为`*-dev.*` `*-pro.*` `*-test.*`三种模式。通过`profile`参数值匹配文件名，实现不同环境配置加载。

## 动态配置加载

配置中心配置可直接解析刷新到配置对象中

```go
var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
mm:
 Jacket: leather
 Trousers: denim
age: 35
eyes: brown
beard: true
`)

v := viper.New()
v.SetConfigType("yaml")
_ = v.ReadConfig(bytes.NewBuffer(yamlExample))
gocfg.WriteConfigMap(v.AllSettings())
```

自定义map，数据库等配置文件可动态加载到配置对象中。

```go
gocfg.WriteConfigMap(map[string]interface{}{
  "test": map[string]interface{}{
    "aaa": "TTET",
  },
})
```

