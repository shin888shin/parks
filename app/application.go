package app

import (
	"github.com/gin-gonic/gin"
)

var (
	r = gin.Default()
)

func StartApplication() {
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	mapUrls()
	// logger.Info("about to start the application...")
	r.Run(":8082") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
