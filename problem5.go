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
	fmt.Println("searching")
	res := 1
	allP := make(map[int]bool)
	allP[0] = true
	for i := 1; i <= 20; i++ {
		println("===== checking next prime =====")
		j := maxPowPrime(i)
		if j == 0 {
			println(i, "is not prime")
			continue
		}
		_, ok := allP[j]
		if !ok {
			print("--------> ", j, " didn't exist and was put <--------\n")
			allP[j] = true
		} else {
			println(j, "was already in the map and was ignored")
		}
	}
	delete(allP, 0)
	fmt.Println("These are the numbers that will be multiplied ", allP)
	for x := range allP {
		res *= x
	}
	fmt.Println(res, "is the answer")
}

func maxPowPrime(x int) int {
	for y := 2; y < x; y++ { // Calculates all factors.
		if (x % y) == 0 {
			return 0
		}
	}
	println(x, "is prime")
	pow := 2
	p := int(math.Pow(float64(x), float64(pow)))
	last := x
	for p < 20 && x > 1 {
		println(x, "powered to", pow, "and is now", p)
		pow++
		last = p
		p = int(math.Pow(float64(x), float64(pow)))
	}
	println(x, "^", pow, "=", p, "was way too big, discarted.")
	return last
}
