package database

import "github.com/coffemanfp/yofiotest/assigners"

// AssignersDB is the database assigners handler.
type AssignersDB interface {
	Create(assigners.Assigner) (assigners.Assigner, error)
	GetStats() (assigners.Stats, error)
}
