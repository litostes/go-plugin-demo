package core

import (
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
    "plugin"
)

type Plugin interface {
    ConfigHooks(config *viper.Viper)
    RouterHooks(router *gin.Engine)
}

// config
// router
// middleware
// session
// cookie
// model

func (e *engine) LoadPlugin(mod string, item interface{}) {

    p, err := plugin.Open(e.Path.Plugin + "/" + mod + ".so")
    if err != nil {
        e.Log.Errorf("[ 插件加载失败 ] : %s", err)
        return
    }

    sym, err := p.Lookup("Plugin")
    if err != nil {
        e.Log.Errorf("[ 插件加载失败 ] : %s", err)
        return
    }

    child, ok := sym.(Plugin)
    if !ok {
        e.Log.Errorf("[ 插件加载失败 ] : %s", err)
        return
    }

    child.ConfigHooks(e.Config)
    child.RouterHooks(e.Router)
}
