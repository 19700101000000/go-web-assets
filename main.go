package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	ginFile, err := os.OpenFile("webserver.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = io.MultiWriter(ginFile)
}

func main() {
	r := gin.Default()

	r.Static("/", "./assets")

	r.Run(":8081")
}
