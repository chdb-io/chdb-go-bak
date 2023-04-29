package main

import (
    "fmt"
    "chdb"
    "flag"
)

func main() {
    query := flag.String("query", "SELECT version()", "Query to execute")
    format := flag.String("format", "CSV", "Query output format")
    flag.Parse()
    result := chdb.Query(string(*query), string(*format))
    fmt.Println("Result:", result)
}
