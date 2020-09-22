package core

import (
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
)

func (e *engine) bootPath() {
    e.Path.App = e.Path.Root + "/app"
    e.Path.Routes = e.Path.Root + "/routes"
    e.Path.Plugin = e.Path.Root + "/plugins"
    e.Path.Runtime = e.Path.Root + "/runtime"
}

func (e *engine) bootConfig() {
    e.Config = viper.New()
    e.Config.SetConfigFile(e.Path.Root + "/config.yaml")
    err := e.Config.ReadInConfig()
    if err != nil {
        e.Log.Errorf("Fatal error config file: %s \n", err)
    }
}

func (e *engine) bootRouter() {
    e.Router = gin.New()
    e.Router.Use(e.LoggerToFile(), gin.Recovery())
}

func (e *engine) bootPlugin(mods map[string]interface{}) {
    for mod, item := range mods {
        e.LoadPlugin(mod, item)
    }
}
