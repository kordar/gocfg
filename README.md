# gocfg 

对`viper`进行二次开发，实现`viper`对象分组操作，整合不同初始化方式实现多种方式配置初始化。

## 初始化

- 默认初始化

提供多种初始化形式供不同场景配置使用

```go
InitConfig("./conf/conf.ini") // 默认分组default
```
注：该方式仅支持`ini`配置文件，传参为文件所在路径。

- 设置扩展名

```go
InitDefaultConfig("./conf/conf.yaml", "yaml")  // 默认分组default
```

- 通过目录加载该目录下所有符合扩展的配置

```go
InitConfigWithDir("default", "./conf", "ini", "yaml") 
```

扫描`./conf`目录下的所有`.ini`，`.yaml`的配置文件，并关联到分组`default`下的`viper`(该对象使用合并方式生成)对象。

- 通过目录扫描子目录并生成配置

```go
InitConfigWithParentDir("./conf", "ini", "yaml")  // 请使用InitConfigWithParentDirG指定group名称
InitConfigWithParentDirG("language", "./conf", "ini", "yaml")  // 分组language

// 目录结构
/**
    language
        - en
            - a.ini
        - zh
            - b.yaml
 */

// 获取方式
gocfg.GetSectionValue("zh", "bb", "language")
```

注：该方式生成的对象为单`viper`对象，该对象第一层key名称为子目录第一级。应用场景多为多语言配置项目


- 扫描目录将子目录第一层作为group生成配置

```go
InitConfigWithSubDir(dir string, ext ...string)  
```

注：该方式生成的配置为多个`viper`对象，`group`值为当前目录子目录第一层目录名称。

- 自定义viper对象生成配置

```go
InitCustomerConfigWithDir(v *viper.Viper, group string, parent string, ext ...string)
```


## 使用

注：所有`API`中`options[0]`参数表示`groupName`值，默认值"`default`"

- 获取`viper`对象

```go
func GetViper(options ...interface{}) *viper.Viper
```

- 设置`viper`对象

```go
func SetViper(v *viper.Viper, options ...interface{})
```

- 向指定`group`中写入字节码

```go
func WriteConfig(b []byte, options ...interface{})
```

- 向指定`group`中写入`map`对象

```go
func WriteConfigMap(cfg map[string]interface{}, options ...interface{})
```

- 更新`key`中的`value`值

```go
func UpdateValue(key string, value interface{}, options ...interface{})
```

- 查询相关`API`

```go
func Get(key string, options ...interface{}) interface{}
func GetSystemValue(key string, options ...interface{}) string
func GetSettingValue(key string, options ...interface{}) string
func GetSectionValue(section string, key string, options ...interface{}) string
func GetSectionValueInt(section string, key string, options ...interface{}) int
func GetSection(section string, options ...interface{}) map[string]string
func UnmarshalKey(section string, raw interface{}, options ...interface{}) error
func Sub(key string, options ...interface{}) *viper.Viper
func AllSections(options ...interface{}) map[string]interface{}
```

## 环境变量扩展

```go
gocfg.InitEnv("G", "profile")
_ = os.Setenv("G_PROFILE", "dev")
```
设置文件名为`*-dev.*` `*-pro.*` `*-test.*`三种模式。通过`profile`参数值匹配文件名，实现不同环境配置加载。