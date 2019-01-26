package controller

import (
	"net/http"
	"webutil/util"
	"mongoweb/model"
	"mongoweb/config"
	"mongoweb/system/database"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"time"
)

func RoleController(r *gin.Engine) {

	var (
		col = config.COLLECTION_USER
	)

	var get2 = func(c *gin.Context) {

		var (
			beans []model.RoleResult
		)

		//------------------------------------------------------------------
		switch GetRoleName(c) {
		case Visitor:
			JSONVistorUnauth(c)
			return

		case User:
			JSONUnauth(c)
			return

		default:
			JSONUnknownRoleUnauth(c)
			return

		case Admin:
		}

		session, err := database.OpenSession()
		if err != nil {
			Log.Warn(err)
			JSONServerError(c)
			return
		}
		defer session.Close()
		collection := session.DB(config.DB).C(col)

		stages := []bson.M{
			{"$group": bson.M{
				"_id": bson.M{
					"roleName":     "$roleName",
					"roleNameShow": "$roleNameShow",
				},
				"count":     bson.M{"$sum": 1},
				"usernames": bson.M{"$push": "$username"},
			}},
		}

		err = collection.Pipe(stages).All(&beans)

		if err != nil {
			Log.Warn(err)
			JSONServerError(c)
			return
		}

		//------------------------------------------------------------------
		c.JSON(http.StatusOK, model.Result{Success: true, Result: beans,})
	}

	var put = func(c *gin.Context) {
		var (
			formbean model.Role
		)
		err := c.ShouldBindJSON(&formbean)
		if err != nil {
			JSONErrorJson(c)
			return
		}

		//------------------------------------------------------------------
		switch GetRoleName(c) {
		case Visitor:
			JSONVistorUnauth(c)
			return

		case User:
			JSONUnauth(c)
			return

		default:
			JSONUnknownRoleUnauth(c)
			return

		case Admin:

		}

		//------------------------------------------------------------------

		session, err := database.OpenSession()
		if err != nil {
			Log.Warn(err)
			JSONServerError(c)
			return
		}
		defer session.Close()
		collection := session.DB(config.DB).C(col)

		timenow := time.Now()

		_, err = collection.UpdateAll(
			bson.M{"roleName": formbean.RoleName},
			bson.M{"$set": bson.M{
				"roleNameShow":  formbean.RoleNameShow,
				"updateTime":    timenow.Format(config.Time2),
				"updateTimeInt": timenow.Unix(),},
			},
		)

		if err != nil {
			Log.Warn(err)
			JSONServerError(c)
			return
		}

		JSONSuccessPutSuccess(c)
	}

	url := util.GetBeanName(model.Role{})

	r.GET(url, get2)
	r.PUT(url, put)

}
