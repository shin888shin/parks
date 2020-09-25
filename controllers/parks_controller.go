package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shin888shin/parks/domain/parks"
	"github.com/shin888shin/parks/services"
	"github.com/shin888shin/parks/utils/errors"
)

func GetPark(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement GetPark")
}

func CreatePark(c *gin.Context) {
	var park parks.Park
	if err := c.ShouldBindJSON(&park); err != nil {
		restErr := errors.NewBadRequestErr("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreatePark(park)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// func SearchParks(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement SearchParks")
// }
