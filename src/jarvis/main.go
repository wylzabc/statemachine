package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func main(){
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	router = gin.Default()

	router.GET("/ping", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/convertto/:name", convertTo)
	router.GET("/get", getCurrentState)


	router.Run()

}




