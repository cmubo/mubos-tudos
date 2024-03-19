package utils

import (
	"strconv"
	"todo/internal/constants"
	"todo/internal/types"
)

func GetPaginationOffset(pagination types.Pagination) int {
	offset := 0

	offset = (pagination.PerPage * pagination.Page) - pagination.PerPage

	return offset
}

func GetPaginationLimit(pagination types.Pagination) int {
	limit := pagination.PerPage

	if limit > constants.PAGINATION_PERPAGE_MAX {
		limit = constants.PAGINATION_PERPAGE_MAX
	}

	if limit == 0 {
		limit = constants.PAGINATION_PERPAGE_DEFAULT
	}

	return limit
}

func GetPaginationQuery(query string, defaultValue int) int {
	if query == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(query)
	if err != nil {
		return defaultValue
	}

	return value
}
