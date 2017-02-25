package store

// Query
type Query map[string]interface{}

// Collection
type Collection interface {
	//All(result interface{}) error
	Find(query Query, result interface{}) error
	// Does not throw exception when multiple matching records exist.
	// Returns first record found.
	FindOne(query Query, result interface{}) error
	Insert(records ...interface{}) error
	// No updates, only inserts
}

// Database
type Database interface {
	Collection(collection string) Collection
}

// Session
type Session interface {
	Close()
	Database(db string) Database
}

// Store
type Store interface {
	Connect() (Session, error)
}

// NewStore gets a datastore. Mongo is default.
func NewStore(url string) Store {
	return newMgoStore(url)
}
