package model

import "github.com/globalsign/mgo/bson"

type SystemLog struct {
	Id            bson.ObjectId `json:"id" bson:"_id"`
	Method        string        `json:"method"`
	Url           string        `json:"url"`
	Json          string        `json:"json"`
	Remoteaddr    string        `json:"remoteaddr"`
	Status        int           `json:"status"`
	UserId        int           `json:"userId"`
	Latency       string        `json:"latency"`
	Duration      int64         `json:"duration"`
	CreateTime    string        `json:"createTime" bson:"createTime"`
	CreateTimeInt int64         `json:"createTimeInt" bson:"createTimeInt"`
}
