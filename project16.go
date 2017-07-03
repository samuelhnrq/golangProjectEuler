package main

import "math/big"

func prob16() {
	a := big.NewInt(2)
	a.Exp(a, big.NewInt(1000), nil)
	ttl := 0
	for _, digit := range a.String() {
		ttl += int(digit - '0')
	}
	println(ttl)
}
