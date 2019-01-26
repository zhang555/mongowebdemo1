package model

type Result struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
	Page    interface{} `json:"page,omitempty"`
}
//
//type ModelTime struct {
//	CreateTime    string `json:"createTime" bson:"createTime"`
//	UpdateTime    string `json:"updateTime" bson:"updateTime"`
//	CreateTimeInt int64  `json:"createTimeInt" bson:"createTimeInt"`
//	UpdateTimeInt int64  `json:"updateTimeInt" bson:"updateTimeInt"`
//}
