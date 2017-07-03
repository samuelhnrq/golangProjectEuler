package main

import (
	"runtime"
	"sync"
)

var (
	hLen    uint
	highest uint
)

func prob14() {
	var wg sync.WaitGroup
	workLoads := uint(runtime.NumCPU())
	wg.Add(int(workLoads))
	unit := 1000000 / workLoads
	for x := uint(1); x <= workLoads; x++ {
		go doWith(unit*(x-1)+1, unit*x, &wg)
	}
	wg.Wait()
	println("The biggest chain is from", highest, "with", hLen, "elements")
}

func doWith(x, y uint, wg *sync.WaitGroup) {
	for ; x < y; x++ {
		var len uint
		for ts := x; ts > 1; len++ {
			if ts%2 == 0 {
				ts = ts / 2
			} else {
				ts = (3 * ts) + 1
			}
			len++
		}
		if len > hLen {
			highest = x
			hLen = len
		}
	}
	wg.Done()
}
