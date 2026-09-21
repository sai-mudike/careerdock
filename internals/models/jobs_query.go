package models

type JobQuery struct {
	SortBy  string
	OrderBy string
	Page    int
	Limit   int
}
