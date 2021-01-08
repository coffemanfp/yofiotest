package api

import (
	"errors"

	"github.com/coffemanfp/yofiotest/database"
	"github.com/coffemanfp/yofiotest/database/postgresql"
)

// DBServices represents the database services.
type DBServices struct {
	Assigners database.AssignersDB
}

// initDBServices starts the database services used by handlers.
func initDBServices(dbS *DBServices) (err error) {
	db := database.Get()
	if db == nil {
		err = errors.New("fatal: non-existent database connection")
		return
	}

	p := db.(*postgresql.PostgreSQL)

	dbS.Assigners = postgresql.NewAssignersDB(p)
	return
}
