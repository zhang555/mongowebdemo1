package controller

import (
	"mongoweb/model"
	"mongoweb/config"
	"time"
	"mongoweb/system/database"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo"
)

func CreateUser(bean *model.User) {

	session, err := database.OpenSession()
	if err != nil {
		Log.Warn(err)
		return
	}
	defer session.Close()
	collection := session.DB(config.DB).C(config.COLLECTION_USER)

	timenow := time.Now()
	bean.Id = bson.NewObjectId()
	bean.CreateTime = timenow.Format(config.Time2)
	bean.CreateTimeInt = timenow.Unix()
	bean.UpdateTime = timenow.Format(config.Time2)
	bean.UpdateTimeInt = timenow.Unix()

	err = collection.Insert(bean)
	if err != nil {
		Log.Warn(err)
		return
	}

}

func InitDB() {

	var (
		result model.User
	)

	session, err := database.OpenSession()
	if err != nil {
		Log.Warn(err)
		return
	}
	defer session.Close()
	collection := session.DB(config.DB).C(config.COLLECTION_USER)

	err = collection.Find(bson.M{"username": "user1"}).One(&result)
	if err == mgo.ErrNotFound {
		user1 := model.User{
			Username:     "user1",
			Password:     "user1",
			RoleName:     "user",
			RoleNameShow: "user",
		}

		CreateUser(&user1)
	} else if err != nil {
		Log.Warn(err)
		return
	}

	err = collection.Find(bson.M{"username": "user2"}).One(&result)
	if err == mgo.ErrNotFound {
		user2 := model.User{
			Username:     "user2",
			Password:     "user2",
			RoleName:     "user",
			RoleNameShow: "user",
		}

		CreateUser(&user2)
	} else if err != nil {
		Log.Warn(err)
		return
	}

	err = collection.Find(bson.M{"username": "admin"}).One(&result)
	if err == mgo.ErrNotFound {
		admin := model.User{
			Username:     "admin",
			Password:     "admin",
			RoleName:     "admin",
			RoleNameShow: "admin",
		}

		CreateUser(&admin)

	} else if err != nil {
		Log.Warn(err)
		return
	}


}
