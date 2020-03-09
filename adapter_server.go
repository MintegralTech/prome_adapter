package main

import (
    "github.com/gin-gonic/gin"
    "./controller"
)

type HandlerFunc func(context *gon.Context) error

func wrapper(handler HandlerFunc) func(context *gin.Context) {
    return func(c *gin.Context) {
        handler(c)
    }
}

func main() {
    router := gin.New()
    router.Use(gin.Recovery())
    router.GET("/helloworld", helloworld)
    router.POST("/receive", controller.Wrapper(controller.Receive))
    //metrics := router.Group("/metrics")
    //{
    //    receive.POST("/aggregate", controller.Wrapper(controller.Receive))
    //}
    router.Run(":1234")
}
