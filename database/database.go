package database

var currentDatabase Database

// Database represents the database connection.
type Database interface {
	Init() error
	Close() error
	Ping() error
}

// Init inits the database handler provided.
func Init(db Database) error {
	return db.Init()
}

// Get gets the current database connection.
func Get() Database {
	return currentDatabase
}
