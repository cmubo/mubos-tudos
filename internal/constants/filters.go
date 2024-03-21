package constants

import (
	"todo/internal/types"
)

var TodoAcceptedFilters = types.AcceptedFilters{
	"completed": {
		Operator: []string{"eq"},
	},
}

var SqlOperators = map[string]string{
	"eq":    "=",
	"noteq": "!=",
	"gt":    ">",
	"gte":   ">=",
	"lt":    "<",
	"lte":   "<=",
}
