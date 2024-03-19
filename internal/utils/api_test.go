package utils

import (
	"fmt"
	"reflect"
	"testing"
	"todo/internal/types"
)

func TestGetSortByMap(t *testing.T) {
	query := "created_at.asc"
	res := GetSortByMap(query, "updated_at.desc")
	expected := types.SortMap{
		"created_at": "ASC",
	}
	eq := reflect.DeepEqual(res, expected)

	if !eq {
		logSortMapComparison(res, expected)
		t.Error("Result was incorrect")
	}

	// Multiple and has a space
	query = "created_at.asc, updated_at.desc"
	res = GetSortByMap(query, "updated_at.desc")
	expected = types.SortMap{
		"created_at": "ASC",
		"updated_at": "DESC",
	}
	eq = reflect.DeepEqual(res, expected)

	if !eq {
		logSortMapComparison(res, expected)
		t.Error("Result was incorrect")
	}

	// Multiple and has a space
	query = "created_at.asc, updated_at.desc"
	res = GetSortByMap(query, "updated_at.desc")
	expected = types.SortMap{
		"created_at": "ASC",
		"updated_at": "DESC",
	}
	eq = reflect.DeepEqual(res, expected)

	if !eq {
		logSortMapComparison(res, expected)
		t.Error("Result was incorrect")
	}

	// wrong formats in order
	query = "created_at.fas, updated_at.desc"
	res = GetSortByMap(query, "updated_at.desc")
	expected = types.SortMap{
		"created_at": "DESC",
		"updated_at": "DESC",
	}
	eq = reflect.DeepEqual(res, expected)

	if !eq {
		logSortMapComparison(res, expected)
		t.Error("Result was incorrect")
	}

	// empty string should return default
	query = ""
	res = GetSortByMap(query, "updated_at.desc")
	expected = types.SortMap{
		"updated_at": "DESC",
	}
	eq = reflect.DeepEqual(res, expected)

	if !eq {
		logSortMapComparison(res, expected)
		t.Error("Result was incorrect")
	}

	// whitespace should return default
	query = "  "
	res = GetSortByMap(query, "updated_at.desc")
	expected = types.SortMap{
		"updated_at": "DESC",
	}
	eq = reflect.DeepEqual(res, expected)

	if !eq {
		logSortMapComparison(res, expected)
		t.Error("Result was incorrect")
	}
}

func logSortMapComparison(x types.SortMap, y types.SortMap) {
	fmt.Println(reflect.TypeOf(x))
	fmt.Println("Response: ")
	for k, v := range x {
		fmt.Printf("Key: %v | value: %v \n", k, v)
	}

	fmt.Println("-----")
	fmt.Println("Expected: ")
	for k, v := range y {
		fmt.Printf("Key: %v | value: %v \n", k, v)
	}
	fmt.Println("End of failed test")
}

func TestGetSortByString(t *testing.T) {
	query := "created_at.asc"
	res := GetSortByString(query, "updated_at DESC")
	expected := "created_at ASC"

	if res != expected {
		t.Errorf("Result was incorrect, got: %v, want: %v.", res, expected)
	}

	query = "created_at.asc,updated_at.asc"
	res = GetSortByString(query, "updated_at DESC")
	expected = "created_at ASC,updated_at ASC"

	if res != expected {
		t.Errorf("Result was incorrect, got: %v, want: %v.", res, expected)
	}

	// empty string should return default
	query = ""
	shouldDefault(t, query, "updated_at DESC")

	// whitespace should return default
	query = "    "
	shouldDefault(t, query, "updated_at DESC")

	// query with whitespace defaults
	query = "updated_at desc"
	shouldDefault(t, query, "updated_at DESC")
	// no direction should default
	query = "updated_at"
	shouldDefault(t, query, "updated_at DESC")

	// no direction should default
	query = "updated_at."
	shouldDefault(t, query, "updated_at DESC")

	// not in accepted sort methods should default
	query = "notarealsort.DESC"
	shouldDefault(t, query, "updated_at DESC")
}

func shouldDefault(t *testing.T, query string, defaultVal string) {
	res := GetSortByString(query, defaultVal)
	expected := defaultVal

	if res != expected {
		t.Errorf("Result was incorrect, got: %v, want: %v.", res, expected)
	}
}
