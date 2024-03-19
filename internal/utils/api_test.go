package utils

import (
	"testing"
)

func TestGetSortByString(t *testing.T) {
	acceptedSortMethods := map[string]string{
		"created_at": "DESC",
		"title":      "DESC",
		"updated_at": "DESC",
	}

	query := "created_at.asc"
	res := GetSortByString(query, "updated_at DESC", acceptedSortMethods)
	expected := "created_at ASC"

	if res != expected {
		t.Errorf("Result was incorrect, got: %v, want: %v.", res, expected)
	}

	query = "created_at.asc,updated_at.asc"
	res = GetSortByString(query, "updated_at DESC", acceptedSortMethods)
	expected = "created_at ASC,updated_at ASC"

	if res != expected {
		t.Errorf("Result was incorrect, got: %v, want: %v.", res, expected)
	}

	// empty string should return default
	query = ""
	shouldDefault(t, query, "updated_at DESC", acceptedSortMethods)

	// whitespace should return default
	query = "    "
	shouldDefault(t, query, "updated_at DESC", acceptedSortMethods)

	// query with whitespace defaults
	query = "created_at desc"
	shouldDefault(t, query, "updated_at DESC", acceptedSortMethods)
	// no direction should default
	query = "updated_at"
	shouldDefault(t, query, "updated_at DESC", acceptedSortMethods)

	// no direction should default
	query = "updated_at."
	shouldDefault(t, query, "updated_at DESC", acceptedSortMethods)

	// not in accepted sort methods should default
	query = "notarealsort.DESC"
	shouldDefault(t, query, "updated_at DESC", acceptedSortMethods)
}

func shouldDefault(t *testing.T, query string, defaultVal string, acceptedSortMethods map[string]string) {
	res := GetSortByString(query, defaultVal, acceptedSortMethods)
	expected := defaultVal

	if res != expected {
		t.Errorf("Result was incorrect, got: %v, want: %v.", res, expected)
	}
}
