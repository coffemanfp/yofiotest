package postgresql

import (
	"database/sql"
	"fmt"

	// Driver
	_ "github.com/lib/pq"
)

var currentConnection *PostgreSQL

// New ...
func New(c Config) (newP *PostgreSQL) {
	newP = &PostgreSQL{
		Config: c,
	}

	currentConnection = newP
	return
}

// Get gets the current database connection.
func Get() *PostgreSQL {
	return currentConnection
}

// PostgreSQL ...
type PostgreSQL struct {
	Config Config

	db *sql.DB
}

// GetConn ...
func (p PostgreSQL) GetConn() *sql.DB {
	return p.db
}

// Init ...
func (p *PostgreSQL) Init() (err error) {
	if p.db != nil {
		err = p.Ping()
		return
	}

	url, err := p.Config.GetURL()
	if err != nil {
		return
	}

	p.db, err = sql.Open("postgres", url)
	if err != nil {
		err = fmt.Errorf("fatal: error opening the postgres connection.\n%v", err)
		return
	}

	err = p.Ping()
	if err != nil {
		return
	}

	currentConnection = p
	return
}

// Ping ...
func (p PostgreSQL) Ping() (err error) {
	err = p.db.Ping()
	if err != nil {
		err = fmt.Errorf("fatal: ping failed.\n%v", err)
	}
	return
}

// Close ...
func (p PostgreSQL) Close() error {
	return p.db.Close()
}
