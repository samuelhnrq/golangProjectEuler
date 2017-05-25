package main

import (
	"math"
)

func sumOfSqr(x int) int {
	ttl := 0
	for i := 1; i <= x; i++ {
		ttl += int(math.Pow(float64(i), 2.0))
	}
	return ttl
}

func sqrOfSum(x int) int {
	ttl := 0
	for i := 1; i <= x; i++ {
		ttl += i
	}
	return int(math.Pow(float64(ttl), 2))
}

func sumSqrDif() {
	f := 100
	f = int(math.Abs(float64(sumOfSqr(f) - sqrOfSum(f))))
	println(f)
}
