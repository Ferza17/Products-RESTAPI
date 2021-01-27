package query

import "strings"

type query struct {
	//Query
	FinalQuery *strings.Builder
}

type Value struct {
}

type QueryBuilder struct {
	Query       *query
	totalValues int
	// Table will be []string if use join methods
	table       string
}

