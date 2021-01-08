package api

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInitRoutes(t *testing.T) {
	assert.Nil(t, initRoutes(gin.Default(), &DBServices{}))
}

func TestFailInitRoutes(t *testing.T) {
	err := initRoutes(gin.Default(), nil)
	assert.Equal(t, "fatal: non-existent database services", err.Error())

	err = initRoutes(nil, &DBServices{})
	assert.Equal(t, "fatal: non-existent engine", err.Error())
}
