package app

import "github.com/shin888shin/parks/controllers"

func mapUrls() {
	r.GET("/ping", controllers.Ping)
	// r.GET("/parks/search", controllers.SearchParks)
	r.POST("/parks", controllers.CreatePark)
	r.GET("/parks/:park_id", controllers.GetPark)
	r.PUT("/parks/:park_id", controllers.UpdatePark)
}
