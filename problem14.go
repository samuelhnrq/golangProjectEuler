package main

import "math"

func prob14() {
	hLen := 0
	highest := 0
	for x := 3; x < 1000000; x++ {
		ts := collatz(x)
		if ts > hLen {
			highest = x
			hLen = ts
		}
	}
	println("\nthe biggest chain is from", highest, "with", hLen, "elements")
}

func rCollatz(x int) []int {
	if x == 1 {
		return []int{1}
	}
	if x%2 == 0 {
		return append([]int{x}, rCollatz(x/2)...)
	}
	return append([]int{x}, rCollatz((3*x)+1)...)
}

func collatz(x int) int {
	ttl := 0
	for x != 1 {
		ttl++
		half := float64(x) / 2
		if half == math.Trunc(half) {
			x = int(half)
		} else {
			x = (3 * x) + 1
		}
	}
	return ttl
}
