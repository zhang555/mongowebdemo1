package controller

import (
	"mongoweb/system/database"
	"mongoweb/config"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo"
)

func FindRoleShowNameByRoleName(roleName string) string {

	var bean struct {
		RoleNameShow string `bson:"roleNameShow"`
	}

	session, err := database.OpenSession()
	if err != nil {
		Log.Warn(err)
		return ""
	}
	defer session.Close()
	collection := session.DB(config.DB).C(config.COLLECTION_USER)

	err = collection.
		Find(bson.M{"roleName": roleName,}, ).
		Select(bson.M{"roleNameShow": roleName,}, ).
		One(&bean)

	if err == mgo.ErrNotFound {
		return ""
	} else if  err != nil{
		Log.Warn(err)
		return ""
	}

	return bean.RoleNameShow
}
