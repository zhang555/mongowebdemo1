package interceptor

import (
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"mongoweb/model"
	"mongoweb/system/database"
	"mongoweb/system/log"
	"mongoweb/config"
	"github.com/globalsign/mgo/bson"
	"github.com/gin-contrib/cors"
	"time"
)

//记录日志
func SaveLog() gin.HandlerFunc {
	return func(c *gin.Context) {

		var systemlog = &model.SystemLog{
			Method:     c.Request.Method,
			Url:        c.Request.URL.String(),
			Remoteaddr: c.Request.RemoteAddr,
		}
		start := time.Now()
		systemlog.CreateTime = start.Format(config.Time2)
		systemlog.CreateTimeInt = start.Unix()

		c.Next()

		systemlog.Id = bson.NewObjectId()
		systemlog.Status = c.Writer.Status()

		end := time.Now()
		latency := end.Sub(start)

		systemlog.Latency = latency.String()
		systemlog.Duration = int64(latency)

		session, err := database.OpenSession()
		if err != nil {
			log.Log.Warn(err)
			return
		}
		defer session.Close()

		collection := session.DB(config.DB).C(config.COLLECTION_SYSTEMLOG)

		err = collection.Insert(&systemlog)
		if err != nil {
			log.Log.Warn(err)
			return
		}
	}
}

//设置session
func SetSession() func(r *gin.Context) {
	store := sessions.NewCookieStore([]byte("mongowebsession;;"))
	store.Options(sessions.Options{Path: "/", MaxAge: 60 * 60 * 24 * 365})
	return sessions.Sessions("mongowebsession", store)
}

//
func CORS() gin.HandlerFunc {

	//允许所有域， 允许带cookie
	corsconfig := cors.Config{
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Access-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}

	return cors.New(corsconfig)

}
