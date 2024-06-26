package gocfg

import "github.com/spf13/viper"

const (
	DEV  = "dev"
	PRO  = "pro"
	TEST = "test"
)

func InitEnv(prefix string, envs ...string) {
	viper.AllowEmptyEnv(true)
	viper.SetEnvPrefix(prefix) // 将自动转为大写
	err := viper.BindEnv(envs...)
	if err != nil {
		panic(err)
	}
}
