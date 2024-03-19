package types

type Pagination struct {
	Page    int
	PerPage int
}

type Response struct {
	Message string
	Status  string
}

type SortMap map[string]string
