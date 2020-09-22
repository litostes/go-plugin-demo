package core

import (
    "github.com/sirupsen/logrus"
    "os"
    "path"
    "time"
)

func (e *engine) Logger() {
    filePath := e.Path.Runtime + "/logs"
    name := "system." + time.Now().Format("2006-01-02") + ".log"

    if err := os.MkdirAll(filePath, 0777); err != nil {
        panic(err.Error())
    }

    fileName := path.Join(filePath, name)
    if _, err := os.Stat(fileName); err != nil {
        if _, err := os.Create(fileName); err != nil {
            panic(err.Error())
        }
    }

    src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
    if err != nil {
        panic(err.Error())
    }

    e.Log = logrus.New()

    e.Log.Out = src

    e.Log.SetLevel(logrus.DebugLevel)

    e.Log.SetFormatter(&logrus.JSONFormatter{
        TimestampFormat: "2006-01-02 15:04:05",
    })
}
