package constants

import "todo/internal/types"

var (
	MSG_RESOURCE_CREATED   = "Resource created succesfully"
	MSG_RESOURCE_RETRIEVED = "Resource succesfully retrieved"
	MSG_RESOURCE_UPDATED   = "Resource succesfully updated"
	MSG_RESOURCE_DELETED   = "Resource succesfully deleted"
	MSG_INTERNAL_ERROR     = "Something went wrong"
)

var (
	RESPONSE_SUCCESFULLY_CREATED   = types.Response{Status: "success", Message: MSG_RESOURCE_CREATED}
	RESPONSE_SUCCESFULLY_RETRIEVED = types.Response{Status: "success", Message: MSG_RESOURCE_RETRIEVED}
	RESPONSE_SUCCESFULLY_UPDATED   = types.Response{Status: "success", Message: MSG_RESOURCE_UPDATED}
	RESPONSE_SUCCESFULLY_DELETED   = types.Response{Status: "success", Message: MSG_RESOURCE_DELETED}
)

var PAGINATION_PERPAGE_DEFAULT = 5
var PAGINATION_PERPAGE_MAX = 50
