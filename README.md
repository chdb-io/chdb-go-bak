<a href="https://chdb.fly.dev" target="_blank">
  <img src="https://user-images.githubusercontent.com/1423657/232511039-480548f7-2e51-4a33-949b-15e0a2a79d9c.png" width=130 />
</a>

[![chDB-go](https://github.com/metrico/chdb-go/actions/workflows/chdb.yml/badge.svg)](https://github.com/metrico/chdb-go/actions/workflows/chdb.yml)

# chdb-go
[chDB](https://github.com/auxten/chdb) go bindings for fun and hacking.

### Do not use this!

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
