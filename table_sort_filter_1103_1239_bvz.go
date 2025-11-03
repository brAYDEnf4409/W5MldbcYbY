// 代码生成时间: 2025-11-03 12:39:45
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "log"
)

// Table represents the structure of a table with sortable and filterable fields.
type Table struct {
    ID      int    `json:"id"`
    Name    string `json:"name"`
    Email   string `json:"email"`
    Age     int    `json:"age"`
}

// TableService handles operations related to table data.
type TableService struct{}

// GetTables retrieves a list of tables sorted and filtered based on query parameters.
func (s *TableService) GetTables(ctx iris.Context) {
    // Retrieve query parameters for sorting and filtering.
    sortField := ctx.URLParamDefault("sort", "id")
    sortOrder := ctx.URLParamDefault("order", "asc")
    filterField := ctx.URLParamDefault("filterField", "")
    filterValue := ctx.URLParamDefault("filterValue", "")

    // Define sample data for demonstration purposes.
    tables := []Table{
        {ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30},
        {ID: 2, Name: "Jane Doe", Email: "jane@example.com", Age: 25},
        {ID: 3, Name: "Jim Beam", Email: "jim@example.com", Age: 40},
    }

    // Sort tables based on the provided sort field and order.
    sortTables(tables, sortField, sortOrder)

    // Filter tables based on the provided filter field and value.
    filterTables(tables, filterField, filterValue)

    // Respond with the sorted and filtered data.
    ctx.JSON(iris.StatusOK, tables)
}

// sortTables sorts the table slice based on the provided field and order.
func sortTables(tables []Table, field string, order string) {
    switch field {
    case "id":
        if order == "desc" {
            sort.Slice(tables, func(i, j int) bool {
                return tables[i].ID > tables[j].ID
            })
        } else {
            sort.Slice(tables, func(i, j int) bool {
                return tables[i].ID < tables[j].ID
            })
        }
    case "name":
        if order == "desc" {
            sort.Slice(tables, func(i, j int) bool {
                return tables[i].Name > tables[j].Name
            })
        } else {
            sort.Slice(tables, func(i, j int) bool {
                return tables[i].Name < tables[j].Name
            })
        }
    case "age":
        if order == "desc" {
            sort.Slice(tables, func(i, j int) bool {
                return tables[i].Age > tables[j].Age
            })
        } else {
            sort.Slice(tables, func(i, j int) bool {
                return tables[i].Age < tables[j].Age
            })
        }
    }
}

// filterTables filters the table slice based on the provided field and value.
func filterTables(tables []Table, field string, value string) {
    switch field {
    case "name":
        tables = filterByName(tables, value)
    case "email":
        tables = filterByEmail(tables, value)
    case "age":
        tables = filterByAge(tables, value)
    }
}

// filterByName filters the table slice by name.
func filterByName(tables []Table, value string) []Table {
    var filtered []Table
    for _, table := range tables {
        if table.Name == value {
            filtered = append(filtered, table)
        }
    }
    return filtered
}

// filterByEmail filters the table slice by email.
func filterByEmail(tables []Table, value string) []Table {
    var filtered []Table
    for _, table := range tables {
        if table.Email == value {
            filtered = append(filtered, table)
        }
    }
    return filtered
}

// filterByAge filters the table slice by age.
func filterByAge(tables []Table, value string) []Table {
    age, err := strconv.Atoi(value)
    if err != nil {
        log.Printf("Error parsing age filter value: %s", err)
        return tables
    }
    var filtered []Table
    for _, table := range tables {
        if table.Age == age {
            filtered = append(filtered, table)
        }
    }
    return filtered
}

func main() {
    app := iris.New()
    // Define a new service for table operations.
    tableService := &TableService{}

    // Register the GetTables endpoint.
    app.Get("/tables", tableService.GetTables)

    // Start the IRIS server.
    log.Fatal(app.Listen(":8080"))
}
