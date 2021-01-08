package api_test

import (
	"testing"

	"github.com/coffemanfp/yofiotest/api"
	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	c := api.DefaultConfig()
	assert.NotEmpty(t, c)
}
