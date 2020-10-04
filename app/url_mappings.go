package app

import (
	"github.com/gin-gonic/gin"
	"github.com/shin888shin/parks/controllers"
)

func MapUrls(r *gin.Engine) {
	r.GET("/ping", controllers.Ping)
	r.GET("/foo", controllers.Foo)
	// r.GET("/parks/search", controllers.SearchParks)
	r.POST("/parks", controllers.CreatePark)
	r.GET("/parks", controllers.GetAllParks)
	r.GET("/parks/:park_id", controllers.GetPark)
	r.PUT("/parks/:park_id", controllers.UpdatePark)
}
