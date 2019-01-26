package controller

import (
	"github.com/gin-gonic/gin"
	"mongoweb/config"
	"mongoweb/model"
	"mongoweb/system/database"
	"net/http"
	"webutil/util"
	"strconv"
	"github.com/globalsign/mgo"
)

func SystemLogController(r *gin.Engine) {

	var (
		col = config.COLLECTION_SYSTEMLOG
	)

	var get2 = func(c *gin.Context) {

		var (
			beans []model.SystemLog
			count int
		)

		num, err1 := strconv.Atoi(c.DefaultQuery("pageNum", PageNumString))
		size, err2 := strconv.Atoi(c.DefaultQuery("pageSize", PageSizeString))

		if util.HaveError(err1, err2, ) {
			JSONErrorUrl(c)
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

		session, err := database.OpenSession()
		if err != nil {
			Log.Warn(err)
			JSONServerError(c)
			return
		}
		defer session.Close()
		collection := session.DB(config.DB).C(col)

		err = collection.Find(nil).Limit(size).Skip((num - 1) * size).Sort("-createTimeInt").All(&beans)
		if err != nil && err != mgo.ErrNotFound {
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
		//------------------------------------------------------------------

		pageMap := map[string]int{
			"pageNum":  num,
			"pageSize": size,
			"count":    count,
		}

		c.JSON(http.StatusOK, model.Result{Success: true, Result: beans, Page: pageMap})
		return
	}

	url := util.GetBeanName(model.SystemLog{})

	r.GET(url, get2)

}
