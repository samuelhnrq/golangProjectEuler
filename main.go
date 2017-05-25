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
	prob := []func(){allPrimes, largestPalindrome, smallestMult,
		sumSqrDif, tenThousandPrime}
	if len(os.Args[1]) >= 1 {
		ex, err := strconv.Atoi(os.Args[1])
		if err == nil {
			prob[ex-1]()
			return
		}
	}
	probName := []string{
		"1 - Biggest possible prime",
		"2 - Largest palindrome of three digit multiplication",
		"3 - Smallest number that can be evenly divided by all 1..20",
		"4 - Sum sqr diff of 1000",
		"5 - 10001st prime number",
	}
	fmt.Printf("%v \nHello what's the problem number? ", probName)
	num, err := strconv.Atoi(readString())

	if err != nil {
		fmt.Println("Invalid number!")
		os.Exit(1)
	}
	num--
	if len(prob) > num {
		t := time.Now().UnixNano()
		prob[num]()
		t = time.Now().UnixNano() - t
		fmt.Println("\nTotal time", t, "nanoseconds")
	} else {
		fmt.Println("Problem number not found")
	}

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
