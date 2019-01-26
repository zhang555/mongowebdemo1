package config

import (
	"os"
	"log"
)

const (
	PageNumString  = "1"
	PageSizeString = "10"
	ZeroString     = "0"
	NullString     = ""

	SessionUserId             = "userId"             //string
	SessionUsername           = "username"           //string
	SessionRolename           = "roleName"           //string
	SessionRolenameshow       = "roleNameShow"       //string
	SessionDepartmentname     = "departmentName"     //string
	SessionDepartmentnameshow = "departmentNameShow" //string

	Visitor = "visitor"
	Admin   = "admin"
	User    = "user"

	Time1 = "2006-01-02"
	Time2 = "2006-01-02 15:04:05"
	Time3 = "2006年01月02日"
	Time4 = "2006年01月02日 15时04分05秒"

	DB                   = "mongoweb"
	COLLECTION_USER      = "user"
	COLLECTION_ARTICLE   = "article"
	COLLECTION_SYSTEMLOG = "systemLog"
)

var (
	MONGODB_ADDR string = os.Getenv("MONGODB_ADDR")
)

func init() {

	addr := os.Getenv("MONGODB_ADDR")
	if addr == "" {
		MONGODB_ADDR = "mongodb://192.168.0.112:27017"
	} else {
		MONGODB_ADDR = addr
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.LstdFlags)

}
