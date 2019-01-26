package model

import "github.com/globalsign/mgo/bson"

type User struct {
	Id                 bson.ObjectId `json:"id" bson:"_id"`                                //
	Username           string        `json:"username,omitempty" bson:"username"`           //
	Password           string        `json:"password" bson:"password"`                     //
	RoleName           string        `json:"roleName" bson:"roleName"`                     //
	RoleNameShow       string        `json:"roleNameShow" bson:"roleNameShow"`             //
	DepartmentName     string        `json:"departmentName" bson:"departmentName"`         //
	DepartmentNameShow string        `json:"departmentNameShow" bson:"departmentNameShow"` //
	CreateTime         string        `json:"createTime" bson:"createTime"`                 //
	UpdateTime         string        `json:"updateTime" bson:"updateTime"`                 //
	CreateTimeInt      int64         `json:"createTimeInt" bson:"createTimeInt"`           //
	UpdateTimeInt      int64         `json:"updateTimeInt" bson:"updateTimeInt"`           //
	Contents           []Content     `json:"contents,omitempty" bson:"contents"`           //
	//DepartmentId   int           `json:"departmentId" bson:"departmentId"`     //
	//RoleId         int           `json:"roleId" bson:"roleId"`                 //
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
