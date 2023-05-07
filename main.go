package main

import (
    "fmt"
    "flag"
    "github.com/chdb-io/chdb-go/chdb"
)

func main() {
    query := flag.String("query", "SELECT version()", "Query to execute")
    format := flag.String("format", "CSV", "Query output format")
    flag.Parse()
    result := chdb.Query(string(*query), string(*format))
    fmt.Println(result)
}
