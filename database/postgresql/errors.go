package postgresql

import "fmt"

func errorPrepare(err error) error {
	return fmt.Errorf("fatal: error to prepare statement.\n%s", err)
}

func errorQuery(err error) error {
	return fmt.Errorf("fatal: error to query statement.\n%s", err)
}
