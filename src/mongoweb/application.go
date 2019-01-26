package main

import (
	"fmt"
	"net/http"
	"mongoweb/config"
	"mongoweb/controller"
	"github.com/gin-gonic/gin"
	"mongoweb/controller/interceptor"
)

func main() {

	//-------------------------------------------------------------------------------
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	//r := gin.Default()
	//-------------------------------------------------------------------------------

	r.Use(interceptor.CORS())
	r.Use(interceptor.SaveLog())
	r.Use(interceptor.SetSession())

	//-------------------------------------------------------------------------------

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "mongowebdemo1")
	})

	//-------------------------------------------------------------------------------

	controller.UserController(r)
	controller.RoleController(r)
	controller.LoginRelateController(r)

	controller.ContentController(r)

	controller.SystemLogController(r)

	//-------------------------------------------------------------------------------

	fmt.Println("config.MONGODB_ADDR:  ", config.MONGODB_ADDR)

	r.Run(":8228")
}
