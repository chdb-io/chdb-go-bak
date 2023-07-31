package chdb

import (
    "testing"
    "strings"
)

var testCases = []struct {
    query       string
    expected    string
    description string
}{
    {"SELECT 1", "1", "Basic Select"},
    {"SELECT 'hello'", "\"hello\"", "Basic Select"},
}

func TestChdb(t *testing.T) {
    for _, test := range testCases {
        observed := Query(test.query, "CSV")
        observed = strings.Replace(observed, "\n", "", -1)
        if observed != test.expected {
            t.Fatalf("%s: %s is not %s", test.query, observed, test.expected)
        }
    }
}
