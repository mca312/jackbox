package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	setupRouter(r)
	r.Run()
}

func setupRouter(r *gin.Engine) {
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}
