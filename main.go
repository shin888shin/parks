package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shin888shin/parks/app"
)

func main() {
	r := setupRouter()
	// app.MapUrls(r)
	r.Run(":8082")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	app.MapUrls(r)
	return r
}
