package api_test

import (
	"testing"

	"github.com/coffemanfp/yofiotest/api"
	"github.com/coffemanfp/yofiotest/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewAPI(t *testing.T) {
	c := api.Config{
		Port:   8080,
		Engine: &gin.Engine{},
		DB:     database.Get(),
	}
	newAPI := api.NewAPI(c)

	assert.NotEmpty(t, newAPI)
	assert.Equal(t, c, newAPI.Config)
}
