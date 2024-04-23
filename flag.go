package gocfg

import (
	"flag"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// type myFlag struct {}
// func (f myFlag) HasChanged() bool { return false }
// func (f myFlag) Name() string { return "my-flag-name" }
// func (f myFlag) ValueString() string { return "my-flag-value" }
// func (f myFlag) ValueType() string { return "string" }

// 一旦你的 flag 实现了这个接口，你可以很方便地告诉Viper绑定它：
// viper.BindFlagValue("my-flag-name", myFlag{})

func BindFlagValues(flags viper.FlagValueSet) {
	err := viper.BindFlagValues(flags)
	if err != nil {
		panic(err)
	}
}

// type myFlagSet struct {
//	flags []myFlag
//}
//
//func (f myFlagSet) VisitAll(fn func(FlagValue)) {
//	for _, flag := range flags {
//		fn(flag)
//	}
//}

// fSet := myFlagSet{
//	flags: []myFlag{myFlag{}, myFlag{}},
//}
// viper.BindFlagValues("my-flags", fSet)

func BindFlagValue(key string, flag viper.FlagValue) {
	err := viper.BindFlagValue(key, flag)
	if err != nil {
		panic(err)
	}
}

type ProfileFlag struct{}

func (f ProfileFlag) HasChanged() bool    { return false }
func (f ProfileFlag) Name() string        { return "profile" }
func (f ProfileFlag) ValueString() string { return "dev" }
func (f ProfileFlag) ValueType() string   { return "string" }

//profile  string
//logLevel string
//confile  string

// // using standard library "flag" package
// flag.Int("flagname", 1234, "help message for flagname")
//
// pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
// pflag.Parse()
// viper.BindPFlags(pflag.CommandLine)
//
// i := viper.GetInt("flagname") // retrieve value from viper

func BindPFlags() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}
}
