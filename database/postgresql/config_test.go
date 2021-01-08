package postgresql_test

import (
	"testing"

	"github.com/coffemanfp/yofiotest/database/postgresql"
	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	assert.NotEmpty(t, postgresql.DefaultConfig())
}

func TestGetURL(t *testing.T) {
	tests := []struct {
		name        string
		config      postgresql.Config
		expectedURL string
	}{
		{
			name: "with config fields",
			config: postgresql.Config{
				Name:     "example",
				User:     "example",
				Password: "example",
				Host:     "localhost",
				Port:     5432,
				SslMode:  "disable",
			},
			expectedURL: "dbname=example host=localhost password=example port=5432 user=example sslmode=disable",
		},
		{
			name: "with url field",
			config: postgresql.Config{
				URL: "postgres://example:example@localhost:5432/example",
			},
			expectedURL: "dbname=example host=localhost password=example port=5432 user=example",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url, err := tt.config.GetURL()
			assert.Nil(t, err)
			assert.Equal(t, tt.expectedURL, url)
		})
	}
}

func TestFailGetURL(t *testing.T) {
	c := postgresql.Config{
		URL: "invalid url",
	}

	url, err := c.GetURL()
	assert.Contains(t, err.Error(), "invalid or not provided db url")
	assert.Empty(t, url)
}
