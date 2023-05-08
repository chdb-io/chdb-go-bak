<a href="https://chdb.fly.dev" target="_blank">
  <img src="https://user-images.githubusercontent.com/1423657/236688026-812c5d02-ddcc-4726-baf8-c7fe804c0046.png" width=130 />
</a>

[![chDB-go](https://github.com/metrico/chdb-go/actions/workflows/chdb.yml/badge.svg)](https://github.com/metrico/chdb-go/actions/workflows/chdb.yml)

# chdb-go
[chDB](https://github.com/auxten/chdb) go bindings for fun and hacking.

### Status

- experimental, unstable, subject to changes
- requires [`libchdb.so`](https://github.com/metrico/libchdb/releases) on the system
- requires `CGO` 

#### Example
```go
import (
    "fmt"
    "github.com/chdb-io/chdb-go/chdb"
)

func main() {
    result := chdb.Query("SELECT version()", "CSV")
    fmt.Println(result)
}
```

<br>

:wave: _C/GO developer? Jump in and help us evolve this prototype into a stable module!_
