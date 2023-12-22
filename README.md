<a href="https://chdb.io/" target="_blank">
  <img src="https://avatars.githubusercontent.com/u/132536224" width=130 />
</a>

# chdb-go
[chDB](https://github.com/auxten/chdb) go bindings for chdb

### Status
- experimental
- requires `CGO` 

### Install
1. download latest lib file from [chdb.io](https://github.com/chdb-io/chdb/releases/) which matchs your system
2. put it in ./chdb path
3. run: `go run main.go`, enjoy your go binding with chdb !

#### Example
```go
package main

import (
    "fmt"
    "github.com/chdb-io/chdb-go/chdb"
)

func main() {
    // Stateless Query (ephemeral)
    result := chdb.Query("SELECT version()", "CSV")
    fmt.Println(result)

    // Stateful Query (persistent)
    chdb.Session("CREATE FUNCTION IF NOT EXISTS hello AS () -> 'chDB'", "CSV", "/tmp")
    hello := chdb.Session("SELECT hello()", "CSV", "/tmp")
    fmt.Println(hello)
}
```

<br>
