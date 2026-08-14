package main

import "fmt"

func qpow(base int64, exp int64) int64 {
	var ans int64
	ans = 1
	for exp > 0 {
		if exp%2 == 1 {
			exp = exp - 1
			ans = ans * base
		} else {
			base = base * base
			exp = exp / 2
		}
	}
	return ans
}

func main() {
	var base int64
	var exp int64
	base = 2
	exp = 11
	fmt.Println(qpow(base, exp))
}
