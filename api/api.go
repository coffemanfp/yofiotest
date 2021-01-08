package api

import (
	"errors"
	"fmt"
)

// NewAPI returns a new API instance.
func NewAPI(c Config) API {
	return API{
		Config: c,
	}
}

// API api representation.
type API struct {
	Config     Config
	DBServices DBServices
}

// Run runs the api listener.
func (api API) Run() (err error) {
	err = initDBServices(&api.DBServices)
	if err != nil {
		return
	}
	initRoutes(api.Config.Engine, &api.DBServices)

	if api.Config.Port == 0 {
		err = errors.New("fatal: invalid or not provided: port")
		return
	}

	err = api.Config.Engine.Run(fmt.Sprintf(":%d", api.Config.Port))
	return
}
