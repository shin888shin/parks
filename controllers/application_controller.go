package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping for health status check returns 200
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
