package main

import (
    "github.com/lonely345/go-plugin-demo/core"
    "github.com/lonely345/go-plugin-demo/routes"
    "os"
)

func main() {
    root, _ := os.Getwd()

    core.Engine.Path.Root = root

    core.Engine.Boot()

    routes.InitRouter()

    core.Engine.RunServer()
}
