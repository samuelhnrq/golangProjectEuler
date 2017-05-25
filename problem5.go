package main

import (
	"fmt"
	"math"
)

func smallestMultAlt() {
	println("Starting calculations")
	for x := 21; x < math.MaxInt64; x++ {
		for i := 2; i <= 20; i++ {
			if (x % i) != 0 {
				break
			}
			if i >= 20 {
				println(x, "is the smallest")
				return
			}
		}
	}
}

func smallestMult() {
	fmt.Println("searching for ans.")
	res := 1
	for i := 1; i <= 20; i++ {
		j := maxPowPrime(i)
		if j == 0 {
			continue
		}
		res *= j
	}
	fmt.Println(res, "is the answer")
}

func maxPowPrime(x int) int {
	//checks if prime in the first place
	for y := 2; y < x; y++ {
		if (x % y) == 0 {
			return 0
		}
	}
	exp := 2
	p := x
	tmp := x
	for p < 20 && x > 1 {
		p = int(math.Pow(float64(x), float64(exp)))
		if p > 20 {
			break
		}
		exp++
		tmp = p
	}
	return tmp
}
