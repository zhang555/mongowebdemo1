package controller

import (
	"github.com/globalsign/mgo"
	"github.com/Sirupsen/logrus"
	"mongoweb/system/log"
)

var (
	Session *mgo.Session
	Log     *logrus.Logger
)

func init() {

	Log = log.Log

	InitDB()
}
