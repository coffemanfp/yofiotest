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

	assignersService := api.DBServices.Assigners
	e.POST("/create-assignment", handlers.CreateAssignment(assignersService))
	e.POST("/statistics", handlers.GetStats(assignersService))
}
