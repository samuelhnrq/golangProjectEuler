package main

import "fmt"
import "math"

func tenThousandPrime() {
	nthPrime := 1
	currNum := 2
	currDivisor := 2
	for nthPrime < 10001 {
		stop := int(math.Ceil(math.Sqrt(float64(currNum))))
		for currDivisor < currNum {
			if (currNum % currDivisor) == 0 {
				break
			} else {
				currDivisor++
			}
			if currDivisor > stop {
				nthPrime++
				break
			}
		}
		currNum++
		currDivisor = 2
	}
	fmt.Printf("The largest prime factor is: %d", currNum-1)
}
