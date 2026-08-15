package main

import "fmt"

func famous(base int64, exp int64, mod int64) int64 {
	exp = exp % (mod - 1)
	return (qpow(base, exp, mod))
}
func qpow(base int64, exp int64, mod int64) int64 {
	var ans int64
	ans = 1
	for exp > 0 {
		if exp%2 == 1 {
			exp--
			ans = ans * base % mod
		} else {
			base = base * base % mod
			exp = exp / 2
		}
	}
	return ans

}

func main() {
	var base int64
	var exp int64
	base = 3
	exp = 9223372036854775807
	var mod int64
	mod = 998244353
	fmt.Println(famous(base, exp, mod))
}
