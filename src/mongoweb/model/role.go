package model

type Role struct {
	RoleName     string `json:"roleName" bson:"roleName"`         //
	RoleNameShow string `json:"roleNameShow" bson:"roleNameShow"` //
}

//type RoleResult struct {
//	RoleName      string   `json:"roleName" bson:"_id"`           //
//	RoleNameShows []string `json:"roleNameShows" bson:"roleNameShows"` //
//	Count         int      `json:"count" bson:"count"`                 //
//	Usernames     []string `json:"usernames" bson:"usernames"`         //
//}

type RoleResult struct {
	Role      Role     `json:"role" bson:"_id"`            //
	Count     int      `json:"count" bson:"count"`         //
	Usernames []string `json:"usernames" bson:"usernames"` //
}
