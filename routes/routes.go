package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/lonely345/go-plugin-demo/core"
)

func InitRouter() {
    router := core.Engine.Router.Group("/plugins")
    {
        router.GET("/", func(c *gin.Context) {
            c.JSON(200, core.Engine.Config.GetStringMap("plugin"))
        })
    }
}
