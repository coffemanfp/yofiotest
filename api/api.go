package api

import (
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
	err = api.initDBServices()
	if err != nil {
		return
	}
	api.initRoutes()

	err = api.Config.Engine.Run(fmt.Sprintf(":%d", api.Config.Port))
	return
}
