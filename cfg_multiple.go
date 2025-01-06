package gocfg

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type MultipleCfg struct {
	name  string
	items map[string]*Snippet
}

func NewMultipleCfg(name string) Cfg {
	return &MultipleCfg{
		name:  name,
		items: map[string]*Snippet{},
	}
}

func (l *MultipleCfg) Name() string {
	return l.name
}

func (l *MultipleCfg) group(args ...interface{}) string {
	switch args[0].(type) {
	case string:
		return args[0].(string)
	default:
		return ""
	}
}

func (l *MultipleCfg) getSnippet(args ...interface{}) *Snippet {
	group := l.group(args...)
	return l.items[group]
}

func (l *MultipleCfg) Viper(args ...interface{}) *viper.Viper {
	snippet := l.getSnippet(args...)
	if snippet == nil {
		return nil
	}
	return snippet.GetValue()
}

func (l *MultipleCfg) SetViper(v *viper.Viper, args ...interface{}) {
	group := l.group(args...)
	snippet := NewSnippet(group, v)
	l.items[group] = snippet
}

func (l *MultipleCfg) Write(b []byte, args ...interface{}) {
	snippet := l.getSnippet(args...)
	if snippet != nil {
		snippet.Write(b)
	}
}

func (l *MultipleCfg) WriteMap(cfg map[string]interface{}, args ...interface{}) {
	snippet := l.getSnippet(args...)
	if snippet != nil {
		snippet.WriteMap(cfg)
	}
}

func (l *MultipleCfg) Update(key string, value interface{}, args ...interface{}) {
	snippet := l.getSnippet(args...)
	if snippet != nil {
		snippet.Update(key, value)
	}
}

func (l *MultipleCfg) Get(key string, args ...interface{}) interface{} {
	snippet := l.getSnippet(args...)
	if snippet != nil {
		return snippet.GetValue().Get(key)
	}
	return nil
}

func (l *MultipleCfg) GetSystemValue(key string, args ...interface{}) string {
	snippet := l.getSnippet(args...)
	if snippet != nil {
		return snippet.GetValue().GetString("system." + key)
	}
	return ""
}

func (l *MultipleCfg) GetSettingValue(key string, args ...interface{}) string {
	snippet := l.getSnippet(args...)
	if snippet != nil {
		return snippet.GetValue().GetString("setting." + key)
	}
	return ""
}

func (l *MultipleCfg) GetSectionValue(section string, key string, args ...interface{}) string {
	snippet := l.getSnippet(args...)
	if snippet != nil {
		return snippet.GetValue().GetString(section + "." + key)
	}
	return ""
}

func (l *MultipleCfg) GetSectionValueInt(section string, key string, args ...interface{}) int {
	snippet := l.getSnippet(args...)
	if snippet != nil {
		return snippet.GetValue().GetInt(section + "." + key)
	}
	return 0
}

func (l *MultipleCfg) GetSection(section string, args ...interface{}) map[string]string {
	snippet := l.getSnippet(args...)
	if snippet != nil {
		return snippet.GetValue().GetStringMapString(section)
	}
	return map[string]string{}
}

func (l *MultipleCfg) UnmarshalKey(section string, raw interface{}, args ...interface{}) error {
	snippet := l.getSnippet(args...)
	if snippet != nil {
		return snippet.GetValue().UnmarshalKey(section, raw)
	}
	return errors.Errorf("[gocfg] snippet does not exist.")
}

func (l *MultipleCfg) Sub(key string, args ...interface{}) *viper.Viper {
	snippet := l.getSnippet(args...)
	if snippet != nil {
		return snippet.GetValue().Sub(key)
	}
	return nil
}

func (l *MultipleCfg) AllSettings(args ...interface{}) map[string]interface{} {
	snippet := l.getSnippet(args...)
	if snippet != nil {
		return snippet.GetValue().AllSettings()
	}
	return map[string]interface{}{}
}
