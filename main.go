package main

import (
    "github.com/hbeimf/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello World")
    })

    r.Run("localhost:12312")
}
