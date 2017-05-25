package main

import "fmt"
import "math"

func tenThousandPrime() {
	nthPrime := 1
	currNum := 3
	currDivisor := 2
	for nthPrime < 10001 {
		stop := int(math.Ceil(math.Sqrt(float64(currNum))))
		for currDivisor <= stop {
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
		currNum += 2
		currDivisor = 2
	}
	fmt.Printf("the 10001st prime number is: %d", currNum-2)
}
