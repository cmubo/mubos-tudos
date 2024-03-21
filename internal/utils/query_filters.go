package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"todo/internal/constants"
	"todo/internal/types"

	sq "github.com/Masterminds/squirrel"
)

func GetFiltersFromQuery(query string, acceptedFilters types.AcceptedFilters) []types.Filter {
	var filters = []types.Filter{}

	// if query is empty, return empty filters
	if strings.ReplaceAll(query, " ", "") == "" {
		return filters
	}

	splitQuery := strings.Split(query, "&")

	if len(splitQuery) < 1 {
		return filters
	}

	filters = GetFiltersFromQueries(splitQuery)
	filters = CreateAcceptedFiltersList(filters, acceptedFilters)

	return filters
}

func CreateAcceptedFiltersList(filters []types.Filter, acceptedFilters types.AcceptedFilters) []types.Filter {
	result := []types.Filter{}

	for _, filter := range filters {
		if ok := IsAcceptedFilter(filter, acceptedFilters); ok {
			result = append(result, filter)
		}
	}

	return result
}

func IsAcceptedFilter(filter types.Filter, acceptedFilters types.AcceptedFilters) bool {
	// is the name part of the accepted filters list/map
	acceptedFilter, ok := acceptedFilters[filter.Name]
	if !ok {
		return false
	}

	// The name is part of the accepted filters, is the operator one of the accepter operators. Return since this is the final check
	return IsAcceptedFilterOperator(filter.Operator, acceptedFilter.Operator)
}

func IsAcceptedFilterOperator(operator string, acceptedOperators []string) bool {
	return contains(acceptedOperators, operator)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GetFiltersFromQueries(queries []string) []types.Filter {
	filters := []types.Filter{}

	for _, query := range queries {
		filter, err := GetFilterFromQuery(query)
		if err != nil {
			continue
		}
		filters = append(filters, filter)
	}

	return filters
}

func GetFilterFromQuery(query string) (types.Filter, error) {
	splitQuery := strings.Split(query, "=")
	if len(splitQuery) != 2 {
		return types.Filter{}, errors.New("no value in filter")
	}

	if splitQuery[0] == "" || splitQuery[1] == "" {
		return types.Filter{}, errors.New("invalid filter format")
	}

	value := strings.TrimSpace(splitQuery[1])
	leftSideOfQuery := splitQuery[0]

	name, err := getFilterName(leftSideOfQuery)
	if err != nil {
		return types.Filter{}, err
	}

	operator, err := getFilterOperator(leftSideOfQuery)
	if err != nil {
		return types.Filter{}, err
	}

	return types.Filter{
		Name:     name,
		Operator: operator,
		Value:    value,
	}, nil
}

func getFilterName(v string) (string, error) {
	splitByBracket := strings.Split(v, "[")
	if len(splitByBracket) != 2 {
		return "", errors.New("no value in filter")
	}

	return splitByBracket[0], nil
}

func getFilterOperator(v string) (string, error) {
	// Match content in brackets
	re := regexp.MustCompile(`\[(.*)\]`)
	substr := re.FindString(v)
	if substr == "" {
		return "", errors.New("cant match for operator")
	}

	// Since the match includes the brackets the submatch will grab that first, but also the content within the capture group (.*). We remove the brackets
	str := strings.Trim(substr, "[]")

	return str, nil
}

func AddFiltersToSquirrelQuery(query sq.SelectBuilder, filters []types.Filter) sq.SelectBuilder {
	if len(filters) > 0 {
		for _, filter := range filters {
			operator, ok := constants.SqlOperators[filter.Operator]
			if !ok {
				continue
			}

			// name and operator are approved.
			queryString := fmt.Sprintf("%s %s ?", filter.Name, operator)
			query = query.Where(queryString, filter.Value)
		}
	}

	return query
}
