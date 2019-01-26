package log

import "github.com/Sirupsen/logrus"

var (
	Log *logrus.Logger
)

func init() {
	Log = logrus.New()
	Log.SetFormatter(&logrus.JSONFormatter{})
	//Log.set

}
