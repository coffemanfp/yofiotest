package postgresql

import (
	"fmt"

	"github.com/lib/pq"
)

// DefaultConfig ...
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

// Config ...
type Config struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     int
	SslMode  string
	URL      string
}

// GetURL ...
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
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		c.User,
		c.Password,
		c.Name,
		c.Host,
		c.Port,
		c.SslMode,
	)
	return
}
