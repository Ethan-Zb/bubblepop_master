package api

import (
	"bubblepop_master/api/prop_api"
	"bubblepop_master/api/user_api"
)

type Group struct {
	PropApi prop_api.PropApi
	UserApi user_api.UserApi
}

var GroupApp = new(Group)
