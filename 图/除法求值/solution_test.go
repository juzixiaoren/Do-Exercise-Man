package main

import "testing"

func Test_calcEquation(t *testing.T) {
	equations := [][]string{{"a", "b"}}
	values := []float64{2.0}
	queries := [][]string{{"a", "c"}}
	r := calcEquation(equations, values, queries)
	t.Logf("result: %v", r)
}
