package main

import "testing"

func Test_snakesAndLadders(t *testing.T) {
	ans := snakesAndLadders([][]int{{-1, -1}, {-1, 3}})
	t.Logf("%d", ans)
}
