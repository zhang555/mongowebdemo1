package controller

import (
	"mongoweb/config"
	"mongoweb/controller/session"
)

var (
	GetRoleName = session.GetRoleName
	GetUserId   = session.GetUserId
)

const (
	PageNumString      = config.PageNumString
	PageSizeString     = config.PageSizeString
	ZeroString         = config.ZeroString
	NullString         = config.NullString

	//
	Visitor = config.Visitor
	Admin   = config.Admin
	User    = config.User
)
