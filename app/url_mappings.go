package app

import "github.com/shin888shin/parks/controllers"

func mapUrls() {
	r.GET("/ping", controllers.Ping)
}
