package main

import (
	"math"
)

func sumPrimesBelow2M() {
	total := 2
	currNum := 3
	currDivisor := 0
	foundPrimes := make([]int, 0)
	foundPrimes = append(foundPrimes, 2)
	for currNum < 2000000 {
		ceil := int(math.Ceil(math.Sqrt(float64(currNum))))
		for ; currDivisor < len(foundPrimes); currDivisor++ {
			div := foundPrimes[currDivisor]
			if (currNum % div) == 0 {
				break
			}
			if currDivisor+1 >= len(foundPrimes) || div > ceil {
				foundPrimes = append(foundPrimes, currNum)
				total += currNum
				break
			}
		}
		currNum += 2
		currDivisor = 0
	}
	println("sum of all primes bellow 2m is: ", total)
}
