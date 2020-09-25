package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// TODO: handle error
		return
	}
	if err := json.Unmarshal(bytes, &park); err != nil {
		// TODO: handle error
		return
	}

	result, saveError := services.CreatePark(park)
	if saveError != nil {
		// TODO: handle create park error
		return
	}
	// spew.Dump
	fmt.Printf("+++> park: %+v\n", park)
	fmt.Printf("+++> error: %+v\n", err)
	fmt.Printf("+++> bytes: %+v\n", string(bytes))
	// c.String(http.StatusNotImplemented, "implement CreateParks")
	c.JSON(http.StatusCreated, result)
}

// func SearchParks(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement SearchParks")
// }
