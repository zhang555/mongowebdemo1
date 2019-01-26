package controller

import (
	"time"
	"net/http"
	"mongoweb/model"
	"mongoweb/config"
	"mongoweb/system/database"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gin-contrib/sessions"
)

func LoginRelateController(r *gin.Engine) {

	loginCh := make(chan bool)

	go func() {
		for <-loginCh {
			time.Sleep(200 * time.Millisecond)
		}
	}()

	var (
		col = config.COLLECTION_USER
	)

	//登陆
	var login = func(c *gin.Context) {
		var (
			formbean model.UserLogin
			findbean model.User
		)

		err := c.ShouldBindJSON(&formbean)
		if err != nil {
			JSONErrorJson(c)
			return
		}

		if formbean.Username == "" || formbean.Password == "" {
			JSONErrorJson(c)
			return
		}

		time.Sleep(500 * time.Millisecond)
		loginCh <- true

		session, err := database.OpenSession()
		if err != nil {
			Log.Warn(err)
			JSONServerError(c)
			return
		}
		defer session.Close()

		collection := session.DB(config.DB).C(col)

		query := bson.M{
			"username": formbean.Username,
			"password": formbean.Password,
		}

		err = collection.Find(query).One(&findbean)
		if err == mgo.ErrNotFound {
			JSONErrorMessage(c, "账号密码错误")
			return
		} else if err != nil {
			Log.Warn(err)
			JSONServerError(c)
			return
		}

		websession := sessions.Default(c)

		//先clear
		websession.Clear()
		websession.Save()

		//这个地方出了问题，找了好久，不能用 bson.ObjectId类型，要用string类型
		websession.Set(config.SessionUserId, findbean.Id.Hex())
		websession.Set(config.SessionUsername, formbean.Username)
		websession.Set(config.SessionRolename, findbean.RoleName)
		websession.Set(config.SessionRolenameshow, findbean.RoleNameShow)
		websession.Set(config.SessionDepartmentname, findbean.DepartmentName)
		websession.Set(config.SessionDepartmentnameshow, findbean.DepartmentNameShow)

		//{"level":"warning","msg":"securecookie: error - caused by: securecookie: error - caused by: gob: type not registered for interface: bs
		//on.ObjectId","time":"2019-01-16T10:55:13+08:00"}
		err = websession.Save()
		if err != nil {
			Log.Warn(err)
			JSONServerError(c)
			return
		}

		result := gin.H{
			config.SessionUserId:             websession.Get(config.SessionUserId),
			config.SessionUsername:           websession.Get(config.SessionUsername),
			config.SessionRolename:           websession.Get(config.SessionRolename),
			config.SessionRolenameshow:       websession.Get(config.SessionRolenameshow),
			config.SessionDepartmentname:     websession.Get(config.SessionDepartmentname),
			config.SessionDepartmentnameshow: websession.Get(config.SessionDepartmentnameshow),
		}

		JSONSuccessMessageAndResult(c, "登陆成功", result)
	}

	//
	var IsLoginController = func(c *gin.Context) {
		session := sessions.Default(c)
		v := session.Get(config.SessionUserId)

		_, isLogin := v.(int)

		c.JSON(http.StatusOK, model.Result{Success: true, Result: isLogin})
	}

	var Logout = func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.JSON(http.StatusOK, model.Result{Success: true, Result: "注销成功"})
	}

	//显示session中的内容
	var Session = func(c *gin.Context) {
		websession := sessions.Default(c)
		result := gin.H{
			config.SessionUserId:       websession.Get(config.SessionUserId),
			config.SessionUsername:     websession.Get(config.SessionUsername),
			config.SessionRolename:     websession.Get(config.SessionRolename),
			config.SessionRolenameshow: websession.Get(config.SessionRolenameshow),
		}
		c.JSON(http.StatusOK, model.Result{Success: true, Result: result})
	}

	r.POST("/login", login)
	r.PUT("/logout", Logout)
	r.GET("/isLogin", IsLoginController)
	r.GET("/session", Session)
}
