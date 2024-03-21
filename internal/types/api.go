package types

type Pagination struct {
	Page    int
	PerPage int
}

type Response struct {
	Message string
	Status  string
}

type Filter struct {
	Name     string
	Operator string
	Value    string
}

type AcceptedFilters map[string]AcceptedFilterValues
type AcceptedFilterValues struct {
	Operator []string
}
