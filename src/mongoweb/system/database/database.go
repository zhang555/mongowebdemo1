package database

import (
	"github.com/globalsign/mgo"
	"mongoweb/config"
)

func OpenSession() (*mgo.Session, error) {
	session, err := mgo.Dial(config.MONGODB_ADDR)
	if err != nil {
		return nil, err
	}
	return session, nil
}
