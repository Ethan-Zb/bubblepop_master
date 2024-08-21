package api

import "bubblepop_master/api/prop_api"

type ApiGroup struct {
	PropApi prop_api.PropApi
}

var ApiGroupApp = new(ApiGroup)
