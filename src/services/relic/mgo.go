package store

import "gopkg.in/mgo.v2"

type mgoStore struct {
	url     string
	session *mgo.Session
}

func (m *mgoStore) Connect() (Session, error) {
	s, err := mgo.Dial(m.url)

	if err != nil {
		return nil, err
	}
	s.SetMode(mgo.Primary, true)
	m.session = s

	return m, nil
}

func (m *mgoStore) Database(db string) Database {
	return &mgoDatabase{
		db: m.session.DB(db),
	}
}

func (m *mgoStore) Close() {
	m.session.Close()
}

type mgoDatabase struct {
	db *mgo.Database
}

func (m *mgoDatabase) Collection(collection string) Collection {
	return &mgoCollection{
		db:         m.db,
		collection: collection,
	}
}

type mgoCollection struct {
	db         *mgo.Database
	collection string
}

func (m *mgoCollection) Find(query Query, result interface{}) error {
	return m.db.C(m.collection).Find(query).All(result)
}

func (m *mgoCollection) FindOne(query Query, result interface{}) error {
	return m.db.C(m.collection).Find(query).One(result)
}

func (m *mgoCollection) Insert(records ...interface{}) error {
	return m.db.C(m.collection).Insert(records...)
}

func newMgoStore(url string) Store {
	return &mgoStore{
		url: url,
	}
}
