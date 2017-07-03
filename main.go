package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

var s = bufio.NewScanner(os.Stdin)

func main() {
	prob := []func(){nil, nil, allPrimes, largestPalindrome, smallestMult,
		sumSqrDif, tenThousandPrime, greatestProductOneThousand, pythaghorasTriplet,
		sumPrimesBelow2M, highestOfGrid, nthTriangleNumber, largeSum, prob14, prob15,
		prob16}
	if len(os.Args) < 1 {
		fmt.Println("Provide the problem number")
	}
	ex, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid number")
	}
	if ex > len(prob) {
		fmt.Println("Problem not implemented yet.")
	}
	t := time.Now().UnixNano()
	prob[ex-1]()
	t = time.Now().UnixNano() - t
	fmt.Printf("Time: %dns %dms\n", t, t/1000000)
	return

}

func readString() string {
	if !s.Scan() {
		log.Fatal("couldn't read from stdio")
	}
	return s.Text()
}

func handleErr(errs ...error) {
	for _, x := range errs {
		if x != nil {
			log.Fatal("Error detected!: ", x.Error())
		}
	}
}

func allPrimes() {
	fmt.Print("Type a number to calculate all prime factors. ")
	x, err := strconv.Atoi(readString())
	handleErr(err)
	y := 2
	for x > 1 {
		if (x % y) == 0 {
			fmt.Println(x, "is dvisible by", y, "added.")
			x = x / y
		} else {
			y++
		}
	}
	fmt.Printf("The largest prime factor is: %d", y)
}

func reverse(s string) bool {
	runes := []rune(s)
	max := int(math.Floor(float64(len(s) / 2.0)))
	for i := 0; i < max; i++ {
		if runes[i] != runes[(len(runes)-1)-i] {
			return false
		}
	}
	return true
}

func largestPalindrome() {
	x, y := 999, 999
	res := ""
	bigger := 0
	for ; x > 100; x-- {
		for ; y > 100; y-- {
			ttl := x * y
			if ttl > bigger {
				res = strconv.Itoa(ttl)
				if reverse(res) {
					bigger = ttl
				}
			} else {
				break
			}
		}
		y = 999
		if x*y < bigger {
			break
		}
	}
	fmt.Print(bigger, " is the largest palindrome of 3 digit multiplication")
}
