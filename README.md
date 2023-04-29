<a href="https://chdb.fly.dev" target="_blank">
  <img src="https://user-images.githubusercontent.com/1423657/232511039-480548f7-2e51-4a33-949b-15e0a2a79d9c.png" width=130 />
</a>

[![chDB-go](https://github.com/metrico/chdb-go/actions/workflows/chdb.yml/badge.svg)](https://github.com/metrico/chdb-go/actions/workflows/chdb.yml)

# chdb-go
[chDB](https://github.com/auxten/chdb) go bindings for fun and hacking.

### Do not use this!

- experimental, unstable, subject to changes
- requires `libchdb.so` on the system _(see releases or action)_

#### Function
```
chdb.Query("SELECT version()", "CSV")
```
```
Query: SELECT version()
Format: CSV
Result: "22.12.1.1"
```
