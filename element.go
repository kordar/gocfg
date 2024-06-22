package gocfg

import (
	"bytes"
	logger "github.com/kordar/gologger"
	"github.com/spf13/viper"
	"sync"
)

type Element struct {
	rw   sync.RWMutex
	name string
	v    *viper.Viper
}

func NewElement(name string) *Element {
	return &Element{name: name}
}

func (g *Element) Name() string {
	return g.name
}

func (g *Element) SetValue(v *viper.Viper) {
	g.rw.Lock()
	g.v = v
	g.rw.Unlock()
}

func (g *Element) Write(b []byte) {
	g.rw.Lock()
	if err := g.v.MergeConfig(bytes.NewBuffer(b)); err != nil {
		logger.Warn("合并配置异常")
	}
	g.rw.Unlock()
}

func (g *Element) WriteMap(m map[string]interface{}) {
	g.rw.Lock()
	if err := g.v.MergeConfigMap(m); err != nil {
		logger.Warn("合并配置异常")
	}
	g.rw.Unlock()
}

func (g *Element) Update(key string, value interface{}) {
	g.rw.Lock()
	g.v.Set(key, value)
	g.rw.Unlock()
}

func (g *Element) GetValue() *viper.Viper {
	g.rw.RLock()
	defer g.rw.RUnlock()
	return g.v
}
