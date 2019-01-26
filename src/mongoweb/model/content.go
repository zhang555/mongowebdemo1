package model

import "github.com/globalsign/mgo/bson"

type Content struct {
	Id            bson.ObjectId `json:"id" bson:"_id"`
	Title         string        `json:"title" bson:"title"`
	Content       string        `json:"content" bson:"content"`
	CreateTime    string        `json:"createTime" bson:"createTime"`
	UpdateTime    string        `json:"updateTime" bson:"updateTime"`
	CreateTimeInt int64         `json:"createTimeInt" bson:"createTimeInt"`
	UpdateTimeInt int64         `json:"updateTimeInt" bson:"updateTimeInt"`
}

//
type UserContent struct {
	Id            bson.ObjectId `json:"id" bson:"_id"`                      //
	Username      string        `json:"username,omitempty" bson:"username"` //
	CreateTime    string        `json:"createTime" bson:"createTime"`       //
	UpdateTime    string        `json:"updateTime" bson:"updateTime"`       //
	CreateTimeInt int64         `json:"createTimeInt" bson:"createTimeInt"` //
	UpdateTimeInt int64         `json:"updateTimeInt" bson:"updateTimeInt"` //
	Content       Content       `json:"content,omitempty" bson:"content"`   //
}

type ContentAndUser struct {
	Id            bson.ObjectId `json:"id" bson:"_id"`
	Title         string        `json:"title" bson:"title"`
	Content       string        `json:"content" bson:"content"`
	CreateTime    string        `json:"createTime" bson:"createTime"`
	UpdateTime    string        `json:"updateTime" bson:"updateTime"`
	CreateTimeInt int64         `json:"createTimeInt" bson:"createTimeInt"`
	UpdateTimeInt int64         `json:"updateTimeInt" bson:"updateTimeInt"`
	//UserId        bson.ObjectId `json:"userId" bson:"_id"`                  //
	Username string `json:"username,omitempty" bson:"username"` //
}
