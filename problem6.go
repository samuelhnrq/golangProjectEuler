package main

import (
	"math"
)

func sumOfSqr(x int, xc chan int) {
	ttl := 0
	for i := 1; i <= x; i++ {
		ttl += int(math.Pow(float64(i), 2.0))
	}
	xc <- ttl
}

func sqrOfSum(x int, xc chan int) {
	ttl := 0
	for i := 1; i <= x; i++ {
		ttl += i
	}
	xc <- int(math.Pow(float64(ttl), 2))
}

func sumSqrDif() {
	f := 100
	xc := make(chan int)
	yc := make(chan int)
	go sumOfSqr(f, xc)
	go sqrOfSum(f, yc)
	f = int(math.Abs(float64(<-xc - <-yc)))
	println("sum of sqr dif", f)
}
