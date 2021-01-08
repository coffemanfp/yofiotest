package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NoRoute handles the not found routes.
func NoRoute(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, newErrorMessage("not found"))
	return
}
