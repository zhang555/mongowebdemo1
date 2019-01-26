package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mongoweb/config"
)

func GetRoleName(c *gin.Context) string {
	session := sessions.Default(c)
	v := session.Get(config.SessionRolename)

	if v2, ok := v.(string); ok {
		return v2
	} else {
		return ""
	}

}

//如果没有登录 调用这个方法 会导致 panic
func GetUserId(c *gin.Context) string {
	session := sessions.Default(c)
	return session.Get(config.SessionUserId).(string)
}

//如果没有登录 调用这个方法 会导致 panic
func GetDepartmentId(c *gin.Context) string {
	session := sessions.Default(c)
	return session.Get(config.SessionDepartmentname).(string)
}
