package main

import (
	"math"
)

func pythaghorasTriplet() {
	a := 1
	b := 1
	c := 1
	sum := 0
	// a to b
	println("calculating:")
	for ; a < 998; a++ {
		for ; b < 998; b++ {
			for ; c < 998; c++ {
				sum = a + b + c
				if sum == 1000 {
					if (pow(a, 2) + pow(b, 2)) == pow(c, 2) {
						println("answer: ", a*b*c)
						return
					}
				}
			}
			c = 1
		}
		b = 1
	}
	// b to c
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
