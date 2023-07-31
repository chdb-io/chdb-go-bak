package main

import (
    "fmt"
    "flag"
    "github.com/chdb-io/chdb-go/chdb"
)

func main() {
    query := flag.String("query", "SELECT version()", "Query to execute")
    format := flag.String("format", "CSV", "Query output format")
    path := flag.String("path", "", "Table persistence path")
    flag.Parse()

    if (path !== ""){
      result := chdb.Session(string(*query), string(*format), string(*path))
      fmt.Println(result)
    } else {
      result := chdb.Query(string(*query), string(*format))
      fmt.Println(result)
    }
}
