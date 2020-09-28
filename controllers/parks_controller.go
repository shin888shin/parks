package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shin888shin/parks/domain/parks"
	"github.com/shin888shin/parks/services"
	"github.com/shin888shin/parks/utils/errors"
)

func GetPark(c *gin.Context) {
	parkID, err := strconv.ParseInt(c.Param("park_id"), 10, 64)
	if err != nil {
		getErr := errors.NewBadRequestErr("park id error")
		c.JSON(getErr.Status, getErr)
		return
	}

	park, getErr := services.GetPark(parkID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, park)
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

func UpdatePark(c *gin.Context) {
	parkID, err := strconv.ParseInt(c.Param("park_id"), 10, 64)
	if err != nil {
		getErr := errors.NewBadRequestErr("park id error")
		c.JSON(getErr.Status, err)
		return
	}

	var park parks.Park
	if err := c.ShouldBindJSON(&park); err != nil {
		restErr := errors.NewBadRequestErr("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	park.ID = parkID
	result, updateErr := services.UpdatePark(park)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

// func SearchParks(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement SearchParks")
// }
