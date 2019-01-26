package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"mongoweb/model"
)

func JSONErrorUrl(c *gin.Context) {
	c.JSON(http.StatusOK, model.Result{Success: false, Code: http.StatusBadRequest, Message: "url格式错误"})
}

func JSONErrorJson(c *gin.Context) {
	c.JSON(http.StatusOK, model.Result{Success: false, Code: http.StatusBadRequest, Message: "json格式错误"})
}

func JSONUnauth(c *gin.Context) {
	c.JSON(http.StatusOK, model.Result{Success: false, Code: http.StatusUnauthorized, Message: "没有权限"})
}
func JSONVistorUnauth(c *gin.Context) {
	c.JSON(http.StatusOK, model.Result{Success: false, Code: 4011, Message: "游客没有权限"})
}
func JSONCanNotOperateOthers(c *gin.Context) {
	c.JSON(http.StatusOK, model.Result{Success: false, Code: http.StatusBadRequest, Message: "不能操作其他人的数据"})
}

func JSONHaveNotLogin(c *gin.Context) {
	c.JSON(http.StatusOK, model.Result{Success: false, Code: http.StatusUnauthorized, Message: "没有登录"})
}
func JSONServerError(c *gin.Context) {
	c.JSON(http.StatusOK, model.Result{Success: false, Code: http.StatusInternalServerError, Message: "服务器内部错误"})
}

//
func JSONErrorMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, model.Result{Success: false, Message: message})
}

func JSONCanNotPutAfterReview(c *gin.Context) {
	c.JSON(http.StatusOK, model.Result{Success: false, Message: "审核后不能修改"})
}

//
func JSONSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, model.Result{Success: true})
}
func JSONSuccessMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, model.Result{Success: true, Message: message})
}

func JSONSuccessMessageAndResult(c *gin.Context, message string, result interface{}) {
	c.JSON(http.StatusOK, model.Result{Success: true, Message: message, Result: result})
}

//
func JSONSuccessPostSuccess(c *gin.Context, id string) {
	c.JSON(http.StatusOK, model.Result{Success: true, Message: "创建成功", Result: id})
}
func JSONSuccessPostSuccessInterface(c *gin.Context, bean interface{}) {
	c.JSON(http.StatusOK, model.Result{Success: true, Message: "创建成功", Result: bean})
}
func JSONSuccessPutSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, model.Result{Success: true, Message: "修改成功"})
}
func JSONSuccessDeleteSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, model.Result{Success: true, Message: "删除成功"})
}
func JSONErrorMessageDBError(c *gin.Context) {
	c.JSON(http.StatusOK, model.Result{Success: false, Message: "操作数据库失败"})
}

//
func JSONSuccessResult(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, model.Result{Success: true, Result: result})
}
func JSONErrorResult(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, model.Result{Success: false, Result: result})
}

func JSONUnknownRoleUnauth(c *gin.Context) {
	c.JSON(http.StatusOK, model.Result{Success: false, Code: http.StatusUnauthorized, Message: "该角色目前没有该权限"})
}
