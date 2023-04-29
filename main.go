package main

import (
    "fmt"
    "chdb"
)

func main() {
    query := "SELECT version()"
    format := "CSV"
    result := chdb.Query("SELECT version()", "CSV")
    fmt.Println("Query:", query)
    fmt.Println("Format:", format)
    fmt.Println("Result:", result)
}
