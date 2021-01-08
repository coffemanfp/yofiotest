package api

import (
	"github.com/coffemanfp/yofiotest/database"
	"github.com/gin-gonic/gin"
)

// DefaultConfig ...
func DefaultConfig() Config {
	return Config{
		Port:   8080,
		Engine: gin.Default(),
		DB:     database.Get(),
	}
}

// Config represents the api settings.
type Config struct {
	Port   int
	Engine *gin.Engine
	DB     database.Database
}
