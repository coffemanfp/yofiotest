package main

import (
	"log"
	"os"
	"strconv"

	"github.com/coffemanfp/yofiotest/api"
	"github.com/coffemanfp/yofiotest/database"
	"github.com/coffemanfp/yofiotest/database/postgresql"
)

var port int

func main() {
	apiConfig := api.DefaultConfig()
	if port != 0 {
		apiConfig.Port = port
	}

	newAPI := api.NewAPI(apiConfig)

	log.Fatalln(newAPI.Run())
}

func init() {
	c := postgresql.DefaultConfig()

	// Environment
	if dbName := os.Getenv("DB_NAME"); dbName != "" {
		c.Name = dbName
	}
	if dbUser := os.Getenv("DB_USER"); dbUser != "" {
		c.User = dbUser
	}
	if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
		c.Password = dbPassword
	}
	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		c.Host = dbHost
	}
	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		var err error
		c.Port, err = strconv.Atoi(dbPort)
		if err != nil {
			log.Fatalf("fatal: invalid or not provided: port.\n%s\n", err)
		}
	}
	if dbSslMode := os.Getenv("DB_SSLMODE"); dbSslMode != "" {
		c.SslMode = dbSslMode
	}
	if dbURL := os.Getenv("DB_URL"); dbURL != "" {
		c.URL = dbURL
	}

	err := database.Init(postgresql.New(c))
	if err != nil {
		log.Fatalf("fatal: fail to start database.\n%s\n", err)
	}
}
