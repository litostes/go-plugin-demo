package core

import (
    "github.com/gin-gonic/gin"
    "time"
)

func (e *engine) LoggerToFile() gin.HandlerFunc {
    return func(c *gin.Context) {
        startTime := time.Now()
        c.Next()
        endTime := time.Now()
        latencyTime := endTime.Sub(startTime)
        reqMethod := c.Request.Method
        reqUri := c.Request.RequestURI
        statusCode := c.Writer.Status()
        clientIp := c.ClientIP()

        e.Log.Infof("| %3d | %13v | %15s | %s | %s |",
            statusCode,
            latencyTime,
            clientIp,
            reqMethod,
            reqUri,
        )
    }
}
