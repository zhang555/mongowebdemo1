package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"webutil/util"
	"mongoweb/model"
	"mongoweb/system/database"
	"net/http"
	"mongoweb/config"
	"github.com/globalsign/mgo/bson"
	"time"
)

func UserController(r *gin.Engine) {

	var (
		col      = config.COLLECTION_USER
		selector = bson.M{
			"username":      1,
			"password":      1,
			"roleName":      1,
			"roleNameShow":  1,
			"createTime":    1,
			"createTimeInt": 1,
			"updateTime":    1,
			"updateTimeInt": 1,
		}
	)

	var get1 = func(c *gin.Context) {
		id := c.Param("id")

		var bean model.User

		//------------------------------------------------------------------
		switch GetRoleName(c) {
		case Visitor:
			JSONVistorUnauth(c)
			return
		case User:
			id = GetUserId(c)

			session, err := database.OpenSession()
			if err != nil {
				Log.Warn(err)
				return
			}
			defer session.Close()
			collection := session.DB(config.DB).C(col)

			err = collection.FindId(bson.ObjectIdHex(id)).Select(selector).One(&bean)
			if err != nil {
				Log.Warn(err)
				JSONServerError(c)
				return
			}

		default:
			JSONUnknownRoleUnauth(c)
			return
		case Admin:

			session, err := database.OpenSession()
			if err != nil {
				Log.Warn(err)
				return
			}
			defer session.Close()
			collection := session.DB(config.DB).C(col)

			err = collection.FindId(bson.ObjectIdHex(id)).Select(selector).One(&bean)
			if err != nil {
				Log.Warn(err)
				JSONServerError(c)
				return
			}

		}
		//------------------------------------------------------------------
		JSONSuccessResult(c, bean)
		return
	}

	var get2 = func(c *gin.Context) {
		num, err1 := strconv.Atoi(c.DefaultQuery("pageNum", PageNumString))
		size, err2 := strconv.Atoi(c.DefaultQuery("pageSize", PageSizeString))

		if util.HaveError(err1, err2, ) {
			JSONErrorUrl(c)
			return
		}

		var (
			beans []model.User
			count int
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

			session, err := database.OpenSession()
			if err != nil {
				Log.Warn(err)
				return
			}
			defer session.Close()
			collection := session.DB(config.DB).C(col)

			err = collection.Find(nil).Select(selector).Limit(size).Skip((num - 1) * size).All(&beans)
			if err != nil {
				Log.Warn(err)
				JSONServerError(c)
				return
			}

			count, err = collection.Find(nil).Count()
			if err != nil {
				Log.Warn(err)
				JSONServerError(c)
				return
			}

		}

		//------------------------------------------------------------------
		pageMap := map[string]int{
			"pageNum":  num,
			"pageSize": size,
			"count":    count,
		}
		c.JSON(http.StatusOK, model.Result{Success: true, Result: beans, Page: pageMap})
	}

	var post = func(c *gin.Context) {
		var (
			bean model.User
		)

		err := c.ShouldBindJSON(&bean)
		if err != nil {
			JSONErrorMessage(c, err.Error())
			return
		}

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

		timenow := time.Now()

		bean.Id = bson.NewObjectId()
		bean.CreateTime = timenow.Format(config.Time2)
		bean.CreateTimeInt = timenow.Unix()
		bean.UpdateTime = timenow.Format(config.Time2)
		bean.UpdateTimeInt = timenow.Unix()

		bean.RoleNameShow = FindRoleShowNameByRoleName(bean.RoleName)

		session, err := database.OpenSession()
		if err != nil {
			Log.Warn(err)
			JSONServerError(c)
			return
		}
		defer session.Close()
		collection := session.DB(config.DB).C(col)

		err = collection.Insert(&bean)
		if err != nil {
			Log.Warn(err)
			JSONServerError(c)
			return
		}

		JSONSuccessPostSuccess(c, bean.Id.Hex())

	}

	var put = func(c *gin.Context) {
		var (
			formbean model.User
			findbean model.User
		)
		err := c.ShouldBindJSON(&formbean)
		if err != nil {
			JSONErrorJson(c)
			return
		}

		if !formbean.Id.Valid() {
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

		err = collection.FindId(bson.ObjectIdHex(formbean.Id.Hex())).One(&findbean)
		if err != nil {
			Log.Warn(err)
			JSONServerError(c)
			return
		}

		findbean.Username = formbean.Username
		findbean.Password = formbean.Password
		findbean.RoleName = formbean.RoleName
		findbean.RoleNameShow = formbean.RoleNameShow

		timenow := time.Now()
		findbean.UpdateTimeInt = timenow.Unix()
		findbean.UpdateTime = timenow.Format(config.Time2)

		err = collection.UpdateId(formbean.Id, &findbean, )
		if err != nil {
			Log.Warn(err)
			JSONServerError(c)
			return
		}

		JSONSuccessPutSuccess(c)
	}

	var dele = func(c *gin.Context) {
		id := c.Param("id")

		var ()
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

		err = collection.RemoveId(bson.ObjectIdHex(id))
		if err != nil {
			Log.Warn(err)
			JSONServerError(c)
			return
		}
		//------------------------------------------------------------------

		JSONSuccessDeleteSuccess(c)

	}

	url := util.GetBeanName(model.User{})

	r.GET(url+"/:id", get1)
	r.GET(url, get2)
	r.POST(url, post)
	r.PUT(url, put)
	r.DELETE(url+"/:id", dele)

}
