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
func (api *API) initDBServices() (err error) {
	p := postgresql.Get()
	if p == nil {
		err = errors.New("fatal: non-existent database connection")
		return
	}

	api.DBServices.Assigners = postgresql.NewAssignersDB(p)
	return
}
