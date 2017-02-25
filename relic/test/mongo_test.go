package test

import (
	"testing"
	"try-hard-platform/relic"
)

type rndStruct struct {
	Name        string
	IsDumb      bool
	notExported string
}

func TestMongoConnection(t *testing.T) {
	// store
	m := store.NewStore("192.168.1.17")

	// session
	session, err := m.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()

	// database
	d := session.Database("tryhard")

	// collection
	c := d.Collection("rndo_collection")

	err = c.Insert(&rndStruct{"rndm", true, "some value"})
	if err != nil {
		t.Fatal(err)
	}
}
