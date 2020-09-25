package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shin888shin/parks/domain/parks"
	"github.com/shin888shin/parks/services"
)

func GetPark(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement GetPark")
}

func CreatePark(c *gin.Context) {
	var park parks.Park
	if err := c.ShouldBindJSON(&park); err != nil {
		// TODO: handle bad request error
		return
	}
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// TODO: handle error
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &park); err != nil {
	// 	// TODO: handle error
	// 	return
	// }

	result, saveError := services.CreatePark(park)
	if saveError != nil {
		// TODO: handle create park error
		return
	}
	c.JSON(http.StatusCreated, result)
}

// func SearchParks(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement SearchParks")
// }
