package query

import (
	"fmt"
	"reflect"
	"strings"
)

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		Query: &query{
			FinalQuery: &strings.Builder{},
		},
	}
}

func (q *QueryBuilder) Insert(table string) *QueryBuilder {
	q.table = table
	q.Query.FinalQuery.WriteString(fmt.Sprintf("INSERT INTO %s", q.table))
	return q
}
func (q *QueryBuilder) Select(table string) *QueryBuilder {
	q.table = table
	q.Query.FinalQuery.WriteString("SELECT ")
	return q
}
func (q *QueryBuilder) Update(table string) *QueryBuilder {
	q.table = table
	q.Query.FinalQuery.WriteString(fmt.Sprintf("UPDATE %s ", q.table))
	return q
}

func (q *QueryBuilder) Set() *QueryBuilder {
	return q
}
func (q *QueryBuilder) Columns(cols interface{}) *QueryBuilder {
	columns := GetColumn(cols)
	q.totalValues = len(columns)

	// SELECT *
	if strings.Contains(q.Query.FinalQuery.String(), "SELECT") && len(columns) == 0 {
		q.Query.FinalQuery.WriteString(fmt.Sprintf("* FROM %s ", q.table))
	} else if strings.Contains(q.Query.FinalQuery.String(), "SELECT") && len(columns) > 0 {
		q.Query.FinalQuery.WriteString("(")
		for i, col := range columns {
			if len(columns) == i {
				q.Query.FinalQuery.WriteString(fmt.Sprintf("%s", col))
				break
			}
			q.Query.FinalQuery.WriteString(fmt.Sprintf("%s,", col))
		}
		q.Query.FinalQuery.WriteString(")")
	}

	// Insert
	if strings.Contains(q.Query.FinalQuery.String(), "INSERT") {
		q.Query.FinalQuery.WriteString("(")
		for i, col := range columns {
			q.Query.FinalQuery.WriteString(fmt.Sprintf("%s", col))
			if i != q.totalValues-1 {
				q.Query.FinalQuery.WriteString(",")
			}
		}
		q.Query.FinalQuery.WriteString(")")
	}

	// Update
	if strings.Contains(q.Query.FinalQuery.String(), "UPDATE") {
		q.Query.FinalQuery.WriteString("SET ")
		for i, col := range columns {

			if len(columns)-1 == i {
				q.Query.FinalQuery.WriteString(fmt.Sprintf("%s=?", col))
				break
			}
			q.Query.FinalQuery.WriteString(fmt.Sprintf("%s=?, ", col))
		}
	}

	return q

}
func (q *QueryBuilder) Where(cols ...string) *QueryBuilder {
	q.Query.FinalQuery.WriteString(" WHERE ")
	for _, col := range cols {
		q.Query.FinalQuery.WriteString(fmt.Sprintf("%s=?", col))
	}

	return q
}
func (q *QueryBuilder) Values() *QueryBuilder {
	q.Query.FinalQuery.WriteString(" VALUES(")

	for i := 0; i < q.totalValues; i++ {
		q.Query.FinalQuery.WriteString("?")
		if i != q.totalValues-1 {
			q.Query.FinalQuery.WriteString(",")
		}
	}
	q.Query.FinalQuery.WriteString(")")

	return q
}

func (q *QueryBuilder) And() *QueryBuilder {
	q.Query.FinalQuery.WriteString(" AND ")
	return q
}

func (q *QueryBuilder) GetValueOf(data interface{}) []interface{} {
	var result []interface{}
	e := reflect.ValueOf(data).Elem()
	for i := 0; i < e.NumField(); i++ {
		varValue := e.Field(i).Interface()

		switch varValue.(type) {
		case string:
			if varValue != "" {
				result = append(result, varValue)
			}
		case int32:
			result = append(result, varValue)
		case int64:
			result = append(result, varValue)
		case float32:
			result = append(result, varValue)
		}
	}
	return RemoveEmptyStringValueOfUser(result, q.Query.FinalQuery.String())
}

func (q *QueryBuilder) BuildQuery() string {
	return q.Query.FinalQuery.String()
}

// Helper
func GetColumn(cols interface{}) []string {
	if cols == nil {
		return []string{}
	}

	e := reflect.ValueOf(cols).Elem()
	str := make([]string, e.NumField())
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varValue := e.Field(i).Interface()

		// TODO: insert to slice if value not nil || not "" || not slice nil
		switch varValue.(type) {
		case string:
			if varValue != "" {
				str = append(str, varName)
			}
		case int32:
			str = append(str, varName)
		case int64:
			str = append(str, varName)
		case float32:
			str = append(str, varName)
		}
	}

	return unique(RemoveEmptyString(str))
}

func unique(strSlice []string) []string {
	keys := make(map[string]bool)
	var list []string

	for _, s := range strSlice {
		if _, value := keys[s]; !value {
			keys[s] = true
			list = append(list, s)
		}
	}
	return list
}

func RemoveEmptyString(strSlice []string) []string {
	var list []string
	for _, s := range strSlice {
		if s != "" {
			list = append(list, s)
		}
	}
	return list
}

func RemoveEmptyStringValueOfUser(strSlice []interface{}, finalQuery string) []interface{} {
	var list []interface{}
	for _, s := range strSlice {
		if s != "" {
			list = append(list, s)
		}
	}
	// insert id
	if strings.Contains(finalQuery, "UPDATE") {
		list = append(list, list[0])
	}
	return list
}
