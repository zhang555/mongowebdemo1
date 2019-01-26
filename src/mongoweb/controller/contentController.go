package controller

import (
	"time"
	"strconv"
	"net/http"
	"webutil/util"
	"mongoweb/model"
	"mongoweb/config"
	"mongoweb/system/database"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo"
)

func ContentController(r *gin.Engine) {

	var (
		col = config.COLLECTION_USER
	)

	var get1 = func(c *gin.Context) {
		id := c.Param("id")

		var bean model.UserContent
		//------------------------------------------------------------------
		switch GetRoleName(c) {
		case Visitor:
			JSONVistorUnauth(c)
			return
		case User:
			userId := GetUserId(c)

			session, err := database.OpenSession()
			if err != nil {
				Log.Warn(err)
				JSONServerError(c)
				return
			}
			defer session.Close()
			collection := session.DB(config.DB).C(col)

			/*
db.user.aggregate(
    [
        {$unwind: "$contents"},
        {
            $match: {_id: ObjectId("5c3d8cea9816112ee4fec5bf"),}
        },
  		{
            $project: {
                content: "$contents",
                username: 1,
            }
        },
		{$sort: {"content.createTimeInt": -1}}
    ]
).pretty()

			*/
			stages := []bson.M{

				{"$unwind": "$contents"},
				{"$match": bson.M{
					"contents._id": bson.ObjectIdHex(id),
					"_id":          bson.ObjectIdHex(userId),
				},},
				{"$project": bson.M{
					"content":  "$contents",
					"username": 1,
				}},
			}

			err = collection.Pipe(stages).One(&bean)
			if err == mgo.ErrNotFound {
				JSONSuccessResult(c, bean)
				return
			} else if err != nil {
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
				JSONServerError(c)
				return
			}
			defer session.Close()
			collection := session.DB(config.DB).C(col)

			stages := []bson.M{

				{"$unwind": "$contents"},
				{"$match": bson.M{
					"contents._id": bson.ObjectIdHex(id),
				},},
				{"$project": bson.M{
					"content":  "$contents",
					"username": 1,
				}},
			}

			err = collection.Pipe(stages).One(&bean)
			if err == mgo.ErrNotFound {
				JSONSuccessResult(c, bean)
				return
			} else if err != nil {
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
			beans     []model.UserContent
			countBean model.CountBean
		)

		session, err := database.OpenSession()
		if err != nil {
			Log.Warn(err)
			JSONServerError(c)
			return
		}
		defer session.Close()
		collection := session.DB(config.DB).C(col)

		//------------------------------------------------------------------
		switch GetRoleName(c) {
		case Visitor:
			JSONVistorUnauth(c)
			return

		case User:

			userId := GetUserId(c)

			/*
db.user.aggregate(
    [
        {$unwind: "$contents"},
        {
            $match: {_id: ObjectId("5c3d8cea9816112ee4fec5bf"),}
        },
  		{
            $project: {
                content: "$contents",
                username: 1,
            }
        },
		{$sort: {"content.createTimeInt": -1}}
    ]
).pretty()
			*/

			m := []bson.M{
				{"$unwind": "$contents"},
				{"$match": bson.M{
					"_id": bson.ObjectIdHex(userId),
				}},
				{"$project": bson.M{
					"content":  "$contents",
					"username": 1,
				}},
				{"$sort": bson.M{"content.createTimeInt": -1}},
				{"$skip": (num - 1) * size},
				{"$limit": size},
			}

			err := collection.Pipe(m).All(&beans)

			//查不到数据竟然返回err， 这个api用起来不爽， 不优雅
			if err == mgo.ErrNotFound {
				pageMap := map[string]int{
					"pageNum":  num,
					"pageSize": size,
					"count":    0,
				}
				c.JSON(http.StatusOK, model.Result{Success: true, Result: beans, Page: pageMap})
				return

			}

			if err != nil {
				Log.Warn(err)
				JSONServerError(c)
				return
			}

			m = []bson.M{
				{"$unwind": "$contents"},
				{"$match": bson.M{
					"_id": bson.ObjectIdHex(userId),
				}},
				{"$count": "count"},
			}

			err = collection.Pipe(m).One(&countBean)
			if err == mgo.ErrNotFound {
				pageMap := map[string]int{
					"pageNum":  num,
					"pageSize": size,
					"count":    0,
				}
				c.JSON(http.StatusOK, model.Result{Success: true, Result: beans, Page: pageMap})
				return

			}
			if err != nil {
				Log.Warn(err)
				JSONServerError(c)
				return
			}

		default:
			JSONUnknownRoleUnauth(c)
			return

		case Admin:

			m := []bson.M{
				{"$unwind": "$contents"},
				{"$project": bson.M{
					"content":  "$contents",
					"username": 1,
				}},
				{"$sort": bson.M{"content.createTimeInt": -1}},

				{"$skip": (num - 1) * size},
				{"$limit": size},
			}

			err := collection.Pipe(m).All(&beans)
			if err == mgo.ErrNotFound {
				pageMap := map[string]int{
					"pageNum":  num,
					"pageSize": size,
					"count":    0,
				}
				c.JSON(http.StatusOK, model.Result{Success: true, Result: beans, Page: pageMap})
				return

			}
			if err != nil {
				Log.Warn(err)
				JSONServerError(c)
				return
			}

			m = []bson.M{
				{"$unwind": "$contents"},
				{"$count": "count"},
			}

			err = collection.Pipe(m).One(&countBean)
			if err == mgo.ErrNotFound {
				pageMap := map[string]int{
					"pageNum":  num,
					"pageSize": size,
					"count":    0,
				}
				c.JSON(http.StatusOK, model.Result{Success: true, Result: beans, Page: pageMap})
				return
			}
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
			"count":    countBean.Count,
		}
		c.JSON(http.StatusOK, model.Result{Success: true, Result: beans, Page: pageMap})
		return

	}

	var post = func(c *gin.Context) {
		var (
			bean model.Content
		)

		err := c.ShouldBindJSON(&bean)
		if err != nil {
			JSONErrorJson(c)
			return
		}

		switch GetRoleName(c) {
		case Visitor:
			JSONVistorUnauth(c)
			return

		case User:

		default:
			JSONUnknownRoleUnauth(c)
			return
		case Admin:

		}

		userId := GetUserId(c)
		session, err := database.OpenSession()
		if err != nil {
			Log.Warn(err)
			JSONServerError(c)
			return
		}
		defer session.Close()
		collection := session.DB(config.DB).C(col)

		timenow := time.Now()

		bean.Id = bson.NewObjectId()
		bean.CreateTime = timenow.Format(config.Time2)
		bean.CreateTimeInt = timenow.Unix()
		bean.UpdateTime = timenow.Format(config.Time2)
		bean.UpdateTimeInt = timenow.Unix()

		//创建时， contents 这个字段名是直接指定的， 不是通过结构体指定的，
		err = collection.UpdateId(
			bson.ObjectIdHex(userId),
			bson.M{"$push": bson.M{"contents": bean},},
		)
		if err != nil {
			Log.Warn(err)
			JSONServerError(c)
			return
		}

		JSONSuccessPostSuccess(c, bean.Id.Hex())
	}

	var put = func(c *gin.Context) {
		var (
			formbean model.Content
			findbean model.UserContent
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
			userId := GetUserId(c)

			session, err := database.OpenSession()
			if err != nil {
				Log.Warn(err)
				JSONServerError(c)
				return
			}
			defer session.Close()
			collection := session.DB(config.DB).C(col)

			stages := []bson.M{
				{"$unwind": "$contents"},
				{"$match": bson.M{
					"contents._id": formbean.Id,
					"_id":          bson.ObjectIdHex(userId),
				},},
				{"$project": bson.M{
					"content":  "$contents",
					"username": 1,
				}},
			}

			err = collection.Pipe(stages).One(&findbean)
			if err == mgo.ErrNotFound {
				JSONCanNotOperateOthers(c)
				return
			} else if err != nil {
				Log.Warn(err)
				JSONServerError(c)
				return
			}

			timenow := time.Now()

			err = collection.Update(
				bson.M{
					"_id":          bson.ObjectIdHex(userId),
					"contents._id": bson.ObjectIdHex(formbean.Id.Hex()),
				},
				bson.M{
					"$set": bson.M{
						"contents.$.title":         formbean.Title,
						"contents.$.content":       formbean.Content,
						"contents.$.updateTime":    timenow.Format(config.Time2),
						"contents.$.updateTimeInt": timenow.Unix(),
					},
				},
			)
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
				JSONServerError(c)
				return
			}
			defer session.Close()
			collection := session.DB(config.DB).C(col)

			timenow := time.Now()

			err = collection.Update(
				bson.M{
					"contents._id": bson.ObjectIdHex(formbean.Id.Hex()),
				},
				bson.M{
					"$set": bson.M{
						"contents.$.title":         formbean.Title,
						"contents.$.content":       formbean.Content,
						"contents.$.updateTime":    timenow.Format(config.Time2),
						"contents.$.updateTimeInt": timenow.Unix(),
					},
				},
			)
			if err != nil {
				Log.Warn(err)
				JSONServerError(c)
				return
			}

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

			/*
"5c401c6f9816110480c2f75c"
5c4034539816112074eb3bc8

指定id，不需要multi
db.user.update(
    {
		_id: 1,
	},
    {
        $pull: {
            contents: {
                _id: ObjectId("5c4034539816112074eb3bc8"),
            }
        }
    }
)

不指定id，需要multi
db.user.update(
    {},
    {
        $pull: {
            contents: {
                _id: ObjectId("5c4034539816112074eb3bc8"),
            }
        }
    },
	{multi: true}
)
			*/

			userId := GetUserId(c)
			session, err := database.OpenSession()
			if err != nil {
				Log.Warn(err)
				JSONServerError(c)
				return
			}
			defer session.Close()
			collection := session.DB(config.DB).C(col)

			err = collection.Update(
				bson.M{
					"_id": bson.ObjectIdHex(userId),
				},
				bson.M{
					"$pull": bson.M{
						"contents": bson.M{"_id": bson.ObjectIdHex(id),
						},
					},
				},
			)

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
				JSONServerError(c)
				return
			}
			defer session.Close()
			collection := session.DB(config.DB).C(col)

			_, err = collection.UpdateAll(
				bson.M{},
				bson.M{
					"$pull": bson.M{
						"contents": bson.M{"_id": bson.ObjectIdHex(id),},
					},
				},
			)

			if err != nil {
				Log.Warn(err)
				JSONServerError(c)
				return
			}
		}

		//------------------------------------------------------------------
		JSONSuccessDeleteSuccess(c)

	}

	url := util.GetBeanName(model.Content{})

	r.GET(url+"/:id", get1)
	r.GET(url, get2)
	r.POST(url, post)
	r.PUT(url, put)
	r.DELETE(url+"/:id", dele)

}
