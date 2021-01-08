package api

import (
	"errors"

	"github.com/coffemanfp/yofiotest/api/handlers"
	"github.com/gin-gonic/gin"
)

func initRoutes(e *gin.Engine, dbS *DBServices) (err error) {
	if dbS == nil {
		err = errors.New("fatal: non-existent database services")
		return
	}
	if e == nil {
		err = errors.New("fatal: non-existent engine")
		return
	}

	e.NoRoute(handlers.NoRoute)

	assignersService := dbS.Assigners
	e.POST("/create-assignment", handlers.CreateAssignment(assignersService))
	e.POST("/statistics", handlers.GetStats(assignersService))
	return
}
