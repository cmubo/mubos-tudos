package utils

import (
	"testing"
	"todo/internal/constants"
	"todo/internal/types"
)

func TestGetPaginationOffset(t *testing.T) {
	offset := GetPaginationOffset(types.Pagination{Page: 1, PerPage: 5})

	if offset != 0 {
		t.Errorf("Result was incorrect, got: %v, want: %v.", offset, 0)
	}

	offset = GetPaginationOffset(types.Pagination{Page: 4, PerPage: 5})

	if offset != 15 {
		t.Errorf("Result was incorrect, got: %v, want: %v.", offset, 15)
	}
}

func TestGetPaginationLimit(t *testing.T) {
	limitMax := constants.PAGINATION_PERPAGE_MAX

	// correct values
	limit := GetPaginationLimit(types.Pagination{Page: 1, PerPage: 7})

	if limit != 7 {
		t.Errorf("Result was incorrect, got: %v, want: %v.", limit, 7)
	}

	// limit past 50 (too large)
	limit = GetPaginationLimit(types.Pagination{Page: 1, PerPage: 100})

	if limit != limitMax {
		t.Errorf("Result was incorrect, got: %v, want: %v.", limit, limitMax)
	}
}

func TestGetPaginationQuery(t *testing.T) {
	defaultValue := 1

	// Empty query
	value := GetPaginationQuery("", 1)

	if value != defaultValue {
		t.Errorf("Result was incorrect, got: %v, want: %v.", value, defaultValue)
	}

	// string not number
	value = GetPaginationQuery("hello", 1)
	if value != defaultValue {
		t.Errorf("Result was incorrect, got: %v, want: %v.", value, defaultValue)
	}

	// correct format
	value = GetPaginationQuery("10", 1)
	if value != 10 {
		t.Errorf("Result was incorrect, got: %v, want: %v.", value, 10)
	}
}
