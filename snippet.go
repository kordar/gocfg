package gocfg

import (
	"bytes"
	logger "github.com/kordar/gologger"
	"github.com/spf13/viper"
	"sync"
)

// Snippet 一般一个文件生成一个snippet对象
type Snippet struct {
	name string
	rw   sync.RWMutex
	v    *viper.Viper
}

func NewSnippet(name string, v *viper.Viper) *Snippet {
	if v == nil {
		v = viper.New()
	}
	return &Snippet{name: name, rw: sync.RWMutex{}, v: v}
}

func (g *Snippet) GetValue() *viper.Viper {
	g.rw.RLock()
	defer g.rw.RUnlock()
	return g.v
}

func (g *Snippet) SetValue(v *viper.Viper) {
	g.rw.Lock()
	g.v = v
	g.rw.Unlock()
}

func (g *Snippet) Name() string {
	return g.name
}

func (g *Snippet) Write(b []byte) {
	g.rw.Lock()
	if err := g.v.MergeConfig(bytes.NewBuffer(b)); err != nil {
		logger.Warn("[gocfg] merge configuration exceptions")
	}
	g.rw.Unlock()
}

func (g *Snippet) WriteMap(m map[string]interface{}) {
	g.rw.Lock()
	if err := g.v.MergeConfigMap(m); err != nil {
		logger.Warn("[gocfg] merge map to configuration exception")
	}
	g.rw.Unlock()
}

func (g *Snippet) Update(key string, value interface{}) {
	g.rw.Lock()
	g.v.Set(key, value)
	g.rw.Unlock()
}
