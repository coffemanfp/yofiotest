package api

import (
	"net/http"

	"github.com/coffemanfp/yofiotest/api/handlers"
	"github.com/gin-gonic/gin"
)

func (api API) initRoutes() {
	e := api.Config.Engine

	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, map[string]string{
			"error": "not found",
		})
		return
	})

	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"message": "Oh yeah!",
		})
	})

	assignersService := api.DBServices.Assigners
	e.POST("/create-assignment", handlers.CreateAssignment(assignersService))
	e.GET("/statistics", handlers.GetStats(assignersService))
}
