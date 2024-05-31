package querybuilder

import (
    "fmt"
    "reflect"
    "strings"
)

type QueryBuilder struct{}

func NewQueryBuilder() *QueryBuilder {
    return &QueryBuilder{}
}

func (qb *QueryBuilder) CreateTable(table interface{}) string {
    t := reflect.TypeOf(table)
    tableName := t.Name()
    var columns []string
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        columns = append(columns, fmt.Sprintf("%s %s", field.Name, field.Type.Name()))
    }
    return fmt.Sprintf("CREATE TABLE %s (%s);", tableName, strings.Join(columns, ", "))
}

func (qb *QueryBuilder) Select(table interface{}) string {
    t := reflect.TypeOf(table)
    tableName := t.Name()
    var columns []string
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        columns = append(columns, field.Name)
    }
    return fmt.Sprintf("SELECT %s FROM %s;", strings.Join(columns, ", "), tableName)
}

func (qb *QueryBuilder) Insert(table interface{}) string {
    t := reflect.TypeOf(table)
    tableName := t.Name()
    var columns []string
    var values []string
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        columns = append(columns, field.Name)
        values = append(values, "?") // Placeholder for prepared statement
    }
    return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tableName, strings.Join(columns, ", "), strings.Join(values, ", "))
}

func (qb *QueryBuilder) Update(table interface{}) string {
    t := reflect.TypeOf(table)
    tableName := t.Name()
    var setClauses []string
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        setClauses = append(setClauses, fmt.Sprintf("%s = ?", field.Name)) // Placeholder for prepared statement
    }
    return fmt.Sprintf("UPDATE %s SET %s WHERE id = ?;", tableName, strings.Join(setClauses, ", "))
}
