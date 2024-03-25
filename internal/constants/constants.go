package constants

import "todo/internal/types"

var (
	MSG_RESOURCE_CREATED         = "Resource created succesfully"
	MSG_RESOURCE_RETRIEVED       = "Resource succesfully retrieved"
	MSG_RESOURCE_UPDATED         = "Resource succesfully updated"
	MSG_RESOURCE_DELETED         = "Resource succesfully deleted"
	MSG_INTERNAL_ERROR           = "Something went wrong"
	MSG_EMAIL_ADDRESS_IN_USE     = "Email address already in use"
	MSG_LOGIN_FAILED_DETAILS     = "Login Failed: Your email or password is incorrect"
	MSG_REGISTRATION_FAILED_BODY = "Registration failed: please check the request body and try again"
	MSG_LOGIN_FAILED_BODY        = "Login failed: please check the request body and try again"
	MSG_FAILED_CREATE_TOKEN      = "Login failed: failed to create a token"
	MSG_LOGIN_SUCCESS            = "Logged in successfully"
	MSG_VALIDATION_SUCCESS       = "Validation successful"
)

var (
	RESPONSE_SUCCESFULLY_CREATED   = types.Response{Status: "success", Message: MSG_RESOURCE_CREATED}
	RESPONSE_SUCCESFULLY_RETRIEVED = types.Response{Status: "success", Message: MSG_RESOURCE_RETRIEVED}
	RESPONSE_SUCCESFULLY_UPDATED   = types.Response{Status: "success", Message: MSG_RESOURCE_UPDATED}
	RESPONSE_SUCCESFULLY_DELETED   = types.Response{Status: "success", Message: MSG_RESOURCE_DELETED}
	RESPONSE_SUCCESFUL_LOGIN       = types.Response{Status: "success", Message: MSG_LOGIN_SUCCESS}
	RESPONSE_SUCCESFUL_VALIDATION  = types.Response{Status: "success", Message: MSG_VALIDATION_SUCCESS}
	RESPONSE_EMAIL_CONFLICT        = types.Response{Status: "error", Message: MSG_EMAIL_ADDRESS_IN_USE}
)

var PAGINATION_PERPAGE_DEFAULT = 5
var PAGINATION_PERPAGE_MAX = 50
