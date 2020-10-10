package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shin888shin/parks/utils/errors"
)

// Ping for health status check returns 200
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pongfriday",
	})
}

// Foo for debugging
func Foo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "barry",
	})
}

// Hostname returns container id
func Hostname(c *gin.Context) {
	name, err := os.Hostname()
	if err != nil {
		restErr := errors.NewBadRequestErr("invalid request")
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"hostname": name,
	})
}
