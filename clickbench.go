package main

import (
    "bufio"
    "fmt"
    "chdb"
    "flag"
    "os"
    "time"
)

func main() {
    _ = chdb.Query("SELECT version()", "CSV")
    sql := flag.String("sql", "queries.sql", "SQL query file")
    flag.Parse()
    readFile, err := os.Open(string(*sql))
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    count := 1
    repeat := 3
    for fileScanner.Scan() {
	for i := 0; i < repeat; i++ {
		start := time.Now()
		_ = chdb.Query(fileScanner.Text(), "CSV")
		elapsed := time.Since(start)
		final := fmt.Sprintf("%d,%d,%.16f", count, i, elapsed.Seconds())
		fmt.Println(final)
	}
	count++
    }
    readFile.Close()
}
