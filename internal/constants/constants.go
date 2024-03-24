package constants

import "todo/internal/types"

var (
	MSG_RESOURCE_CREATED     = "Resource created succesfully"
	MSG_RESOURCE_RETRIEVED   = "Resource succesfully retrieved"
	MSG_RESOURCE_UPDATED     = "Resource succesfully updated"
	MSG_RESOURCE_DELETED     = "Resource succesfully deleted"
	MSG_INTERNAL_ERROR       = "Something went wrong"
	MSG_EMAIL_ADDRESS_IN_USE = "Email address already in use"
)

var (
	RESPONSE_SUCCESFULLY_CREATED   = types.Response{Status: "success", Message: MSG_RESOURCE_CREATED}
	RESPONSE_SUCCESFULLY_RETRIEVED = types.Response{Status: "success", Message: MSG_RESOURCE_RETRIEVED}
	RESPONSE_SUCCESFULLY_UPDATED   = types.Response{Status: "success", Message: MSG_RESOURCE_UPDATED}
	RESPONSE_SUCCESFULLY_DELETED   = types.Response{Status: "success", Message: MSG_RESOURCE_DELETED}
	RESPONSE_EMAIL_CONFLICT        = types.Response{Status: "error", Message: MSG_EMAIL_ADDRESS_IN_USE}
)

var PAGINATION_PERPAGE_DEFAULT = 5
var PAGINATION_PERPAGE_MAX = 50
