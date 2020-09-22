package core

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "golang.org/x/sync/errgroup"
    "log"
    "net/http"
    "os"
    "time"
)

var (
    g errgroup.Group
)

type engine struct {
    Context context
    Path    struct {
        Root    string
        App     string
        Plugin  string
        Routes  string
        Runtime string
    }
    Config *viper.Viper
    Server *http.Server
    Router *gin.Engine
    Log    *logrus.Logger
}

var Engine = &engine{}

func (e *engine) Boot() {
    e.bootPath()
    e.Logger()
    e.bootConfig()
    gin.SetMode(e.Config.GetString("app.mode"))
    e.bootRouter()
    e.bootPlugin(e.Config.GetStringMap("plugin"))
}

func (e *engine) RunServer() {
    e.Server = &http.Server{
        Addr:         e.Config.GetString("server.port"),
        Handler:      e.Router,
        ReadTimeout:  e.Config.GetDuration("server.readTimeout") * time.Second,
        WriteTimeout: e.Config.GetDuration("server.writeTimeout") * time.Second,
    }

    g.Go(func() error {
        return e.Server.ListenAndServe()
    })

    fmt.Printf("Go Http Server Start Successful.\nServer: %s\nPid: %s",
        e.Config.GetString("server.port"),
        fmt.Sprintf("%d", os.Getpid()))

    if err := g.Wait(); err != nil {
        log.Fatal(err)
    }

}
