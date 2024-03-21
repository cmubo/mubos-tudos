package utils

import (
	"reflect"
	"testing"
	"todo/internal/types"
)

var (
	completedFilter = types.Filter{
		Name:     "completed",
		Operator: "eq",
		Value:    "true",
	}

	balanceFilter = types.Filter{
		Name:     "balance",
		Operator: "gte",
		Value:    "4",
	}

	// not accepted
	typeFilter = types.Filter{
		Name:     "type",
		Operator: "eq",
		Value:    "14",
	}

	acceptedFilters = types.AcceptedFilters{
		"completed": {
			Operator: []string{"eq"},
		},
		"balance": {
			Operator: []string{"eq", "neq", "gte", "gt", "lte", "lt"},
		},
	}
)

func TestGetFilterMap(t *testing.T) {
	filters := []string{"completed[eq]=true"}

	res := GetFiltersFromQueries(filters)
	expected := []types.Filter{completedFilter}

	if eq := reflect.DeepEqual(res, expected); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, res)
	}

	filters = []string{"completed[eq]=true", "balance[gte]=4"}
	res = GetFiltersFromQueries(filters)
	expected = []types.Filter{completedFilter, balanceFilter}

	if eq := reflect.DeepEqual(res, expected); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, res)
	}

	// Multiple of same value
	filters = []string{"completed[eq]=true", "balance[gte]=4", "balance[lte]=10"}
	res = GetFiltersFromQueries(filters)
	expected = []types.Filter{
		completedFilter,
		balanceFilter,
		{
			Name:     "balance",
			Operator: "lte",
			Value:    "10",
		},
	}

	if eq := reflect.DeepEqual(res, expected); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, res)
	}

	// Ignore one invalid
	filters = []string{"completed[eq]", "=4"}
	res = GetFiltersFromQueries(filters)
	expected = []types.Filter{}

	if eq := reflect.DeepEqual(res, expected); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, res)
	}

	// ignore one invalid but keep the other
	filters = []string{"completed[eq].true", "balance[gte]=4"}
	res = GetFiltersFromQueries(filters)
	expected = []types.Filter{balanceFilter}

	if eq := reflect.DeepEqual(res, expected); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, res)
	}

	// ignore one invalid (extra equals)
	filters = []string{"completed[eq]=t=rue", "balance[gte]=4"}
	res = GetFiltersFromQueries(filters)
	expected = []types.Filter{balanceFilter}

	if eq := reflect.DeepEqual(res, expected); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, res)
	}
}

func TestIsAcceptedFilter(t *testing.T) {
	// Correct filter
	if ok := IsAcceptedFilter(completedFilter, acceptedFilters); !ok {
		t.Errorf("Result was incorrect expected: %v but got: %v", true, false)
	}

	// Correct filter
	if ok := IsAcceptedFilter(balanceFilter, acceptedFilters); !ok {
		t.Errorf("Result was incorrect expected: %v but got: %v", true, false)
	}

	// Not accepted filter
	if ok := IsAcceptedFilter(typeFilter, acceptedFilters); ok {
		t.Errorf("Result was incorrect expected: %v but got: %v", false, true)
	}

	completedFilterWrongOperator := completedFilter
	completedFilterWrongOperator.Operator = "gte"

	// Not accepted operator
	if ok := IsAcceptedFilter(completedFilterWrongOperator, acceptedFilters); ok {
		t.Errorf("Result was incorrect expected: %v but got: %v", false, true)
	}
}

func TestCreateAcceptedFiltersList(t *testing.T) {
	// Nothing should be filtered out
	filters := []types.Filter{completedFilter, balanceFilter}
	res := CreateAcceptedFiltersList(filters, acceptedFilters)
	if eq := reflect.DeepEqual(res, filters); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", filters, res)
	}

	// Nothing should be filtered out
	filters = []types.Filter{completedFilter}
	res = CreateAcceptedFiltersList(filters, acceptedFilters)
	if eq := reflect.DeepEqual(res, filters); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", filters, res)
	}

	// One should be filtered out
	filters = []types.Filter{completedFilter, typeFilter}
	res = CreateAcceptedFiltersList(filters, acceptedFilters)
	expected := []types.Filter{completedFilter}
	if eq := reflect.DeepEqual(res, expected); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, res)
	}

	// wrong operator on completed filter so it should be filtered out too
	completedFilterCopy := completedFilter
	completedFilterCopy.Operator = "gte"
	filters = []types.Filter{completedFilterCopy, typeFilter}
	res = CreateAcceptedFiltersList(filters, acceptedFilters)
	expected = []types.Filter{}
	if eq := reflect.DeepEqual(res, expected); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, res)
	}
}

func TestGetFiltersFromQuery(t *testing.T) {
	// 1 accepted filter
	query := "completed[eq]=true"
	res := GetFiltersFromQuery(query, acceptedFilters)
	expected := []types.Filter{completedFilter}
	if eq := reflect.DeepEqual(res, expected); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, res)
	}

	// 2 accepted filters
	query = "completed[eq]=true&balance[gte]=4"
	res = GetFiltersFromQuery(query, acceptedFilters)
	expected = []types.Filter{completedFilter, balanceFilter}
	if eq := reflect.DeepEqual(res, expected); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, res)
	}

	// One none accepted filter
	query = "completed[eq]=true&hasProject[eq]=false"
	res = GetFiltersFromQuery(query, acceptedFilters)
	expected = []types.Filter{completedFilter}
	if eq := reflect.DeepEqual(res, expected); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, res)
	}

	// greater than or equal to
	query = "balance[gte]=4"
	res = GetFiltersFromQuery(query, acceptedFilters)
	expected = []types.Filter{balanceFilter}
	if eq := reflect.DeepEqual(res, expected); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, res)
	}

	// less than or equal to
	query = "balance[lte]=30"
	res = GetFiltersFromQuery(query, acceptedFilters)
	expected = []types.Filter{
		{
			Name:     "balance",
			Operator: "lte",
			Value:    "30",
		},
	}
	if eq := reflect.DeepEqual(res, expected); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, res)
	}

	// space around trimmed
	query = "completed[eq]=true"
	res = GetFiltersFromQuery(query, acceptedFilters)
	expected = []types.Filter{completedFilter}
	if eq := reflect.DeepEqual(res, expected); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, res)
	}

	// Blank, ignored
	query = ""
	res = GetFiltersFromQuery(query, acceptedFilters)
	expected = []types.Filter{}
	if eq := reflect.DeepEqual(res, expected); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, res)
	}

	// Empty space, ignored
	query = " "
	res = GetFiltersFromQuery(query, acceptedFilters)
	expected = []types.Filter{}
	if eq := reflect.DeepEqual(res, expected); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, res)
	}

	//  comma seperated, all ignored
	query = "completed[eq]=true,balance[gte]=4"
	res = GetFiltersFromQuery(query, acceptedFilters)
	expected = []types.Filter{}
	if eq := reflect.DeepEqual(res, expected); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, res)
	}

	// Spaces in one, one ignored
	query = "completed[eq]= true&bala nce[gte]=4"
	res = GetFiltersFromQuery(query, acceptedFilters)
	expected = []types.Filter{completedFilter}
	if eq := reflect.DeepEqual(res, expected); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, res)
	}

	// Spaces in both, both ignored
	query = "comple ted[eq]=true&bala nce[gte]=4"
	res = GetFiltersFromQuery(query, acceptedFilters)
	expected = []types.Filter{}
	if eq := reflect.DeepEqual(res, expected); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, res)
	}

	// No value, ignored
	query = "completed[eq]="
	res = GetFiltersFromQuery(query, acceptedFilters)
	expected = []types.Filter{}
	if eq := reflect.DeepEqual(res, expected); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, res)
	}

	// Not in accepted params, ignored
	query = "notAccepted[eq]=true"
	res = GetFiltersFromQuery(query, acceptedFilters)
	expected = []types.Filter{}
	if eq := reflect.DeepEqual(res, expected); !eq {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, res)
	}

}

func TestGetFilterOperator(t *testing.T) {
	str, err := getFilterOperator("completed[eq]")
	expected := "eq"
	if str != expected && err != nil {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, str)
	}

	// Should fail
	str, err = getFilterOperator("completedeq]")
	expected = ""
	if str != expected && err != nil {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, str)
	}

	// Should fail
	str, err = getFilterOperator("completed[eq")
	expected = ""
	if str != expected && err != nil {
		t.Errorf("Result was incorrect expected: %v but got: %v", expected, str)
	}
}
