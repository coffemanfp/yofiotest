package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/coffemanfp/yofiotest/database"
	"github.com/coffemanfp/yofiotest/database/postgresql"
)

var (
	withExamples bool
	examplesFile string
	schemaFile   string
)

var envDatabaseConfig postgresql.Config

func main() {
	pg := postgresql.Get()
	db := pg.GetConn()

	// Get and execute the schema
	schemaBytes, err := getFilebytes(schemaFile)
	if err != nil {
		log.Fatalln("failed to read the schema file: ", err)
	}

	_, err = db.Exec(string(schemaBytes))
	if err != nil {
		log.Fatalln("failed to execute the schema: ", err)
	}

	log.Println("Schema executed successfully!!")

	if !withExamples {
		return
	}

	// Get and execute the examples.
	examplesBytes, err := getFilebytes(examplesFile)
	if err != nil {
		log.Fatalln("failed to read the examples file: ", err)
	}

	_, err = db.Exec(string(examplesBytes))
	if err != nil {
		log.Fatalln("failed to execute the examples: ", err)
	}

	log.Println("Examples executed successfully!!")
}

func init() {
	// Flags
	flag.BoolVar(&withExamples, "with-examples", false, "if sets, reads the examples-file value and creates the examples provided.")
	flag.StringVar(&examplesFile, "examples-file", "./examples.sql", "examples registers file location")
	flag.StringVar(&schemaFile, "schema-file", "./schema.sql", "schema file location")

	flag.Parse()

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

// getFilebytes - Gets the bytes from a file by its filepath.
func getFilebytes(path string) (fileBytes []byte, err error) {
	_, err = os.Stat(path)
	if err != nil {
		return
	}

	fileBytes, err = ioutil.ReadFile(path)
	return
}
