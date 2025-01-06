package gocfg

const (
	DEV  = "dev"
	PRO  = "pro"
	TEST = "test"
)

func InitEnv(prefix string, envs ...string) {
	obj := GetViperObj()
	if obj == nil {
		return
	}
	obj.AllowEmptyEnv(true)
	obj.SetEnvPrefix(prefix) // 将自动转为大写
	err := obj.BindEnv(envs...)
	if err != nil {
		panic(err)
	}
}
