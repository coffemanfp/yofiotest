package postgresql

import (
	"fmt"

	"github.com/lib/pq"
)

// DefaultConfig returns the default database config.
func DefaultConfig() Config {
	return Config{
		Name:     "yofiotest",
		User:     "yofiotest",
		Password: "yofiotest",
		Host:     "localhost",
		Port:     5432,
		SslMode:  "disable",
	}
}

// Config represents the settings.
type Config struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     int
	SslMode  string
	URL      string
}

// GetURL gets the url from the c.URL field or generates the url
// from the other config fields.
func (c Config) GetURL() (url string, err error) {
	if c.URL != "" {
		url, err = pq.ParseURL(c.URL)
		if err != nil {
			err = fmt.Errorf("fatal: invalid or not provided db url.\n%s", err)
			return
		}
		return
	}

	url = fmt.Sprintf(
		"dbname=%s host=%s password=%s port=%d user=%s sslmode=%s",
		c.Name,
		c.Host,
		c.Password,
		c.Port,
		c.User,
		c.SslMode,
	)
	return
}
