package gocfg

import "github.com/spf13/viper"

type LocalCfg struct {
	name    string
	snippet *Snippet
}

func NewLocalCfg(name string, v *viper.Viper) Cfg {
	return &LocalCfg{name: name, snippet: NewSnippet("default", v)}
}

func NewLocalCfgWithSnippet(name string, snippet *Snippet) Cfg {
	return &LocalCfg{name: name, snippet: snippet}
}

func (l *LocalCfg) Name() string {
	return l.name
}

func (l *LocalCfg) Viper(args ...interface{}) *viper.Viper {
	return l.snippet.GetValue()
}

func (l *LocalCfg) SetViper(v *viper.Viper, args ...interface{}) {
	l.snippet.SetValue(v)
}

func (l *LocalCfg) Write(b []byte, args ...interface{}) {
	l.snippet.Write(b)
}

func (l *LocalCfg) WriteMap(cfg map[string]interface{}, args ...interface{}) {
	l.snippet.WriteMap(cfg)
}

func (l *LocalCfg) Update(key string, value interface{}, args ...interface{}) {
	l.snippet.Update(key, value)
}

func (l *LocalCfg) Get(key string, args ...interface{}) interface{} {
	return l.snippet.GetValue().Get(key)
}

func (l *LocalCfg) GetSystemValue(key string, args ...interface{}) string {
	return l.snippet.GetValue().GetString("system." + key)
}

func (l *LocalCfg) GetSettingValue(key string, args ...interface{}) string {
	return l.snippet.GetValue().GetString("setting." + key)
}

func (l *LocalCfg) GetSectionValue(section string, key string, args ...interface{}) string {
	return l.snippet.GetValue().GetString(section + "." + key)
}

func (l *LocalCfg) GetSectionValueInt(section string, key string, args ...interface{}) int {
	return l.snippet.GetValue().GetInt(section + "." + key)
}

func (l *LocalCfg) GetSection(section string, args ...interface{}) map[string]string {
	return l.snippet.GetValue().GetStringMapString(section)
}

func (l *LocalCfg) UnmarshalKey(section string, raw interface{}, args ...interface{}) error {
	return l.snippet.GetValue().UnmarshalKey(section, raw)
}

func (l *LocalCfg) Sub(key string, args ...interface{}) *viper.Viper {
	return l.snippet.GetValue().Sub(key)
}

func (l *LocalCfg) AllSettings(args ...interface{}) map[string]interface{} {
	return l.snippet.GetValue().AllSettings()
}
