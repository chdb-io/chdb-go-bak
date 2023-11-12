<a href="https://chdb.fly.dev" target="_blank">
  <img src="https://avatars.githubusercontent.com/u/132536224" width=130 />
</a>

[![chDB-go](https://github.com/metrico/chdb-go/actions/workflows/chdb.yml/badge.svg)](https://github.com/metrico/chdb-go/actions/workflows/chdb.yml)

# chdb-go
[chDB](https://github.com/auxten/chdb) go bindings for fun and hacking.

### Status

- experimental, unstable, subject to changes
- requires [`libchdb`](https://github.com/metrico/libchdb) on the system
- requires `CGO` 

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

:wave: _C/GO developer? Jump in and help us evolve this prototype into a stable module!_
