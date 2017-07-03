package main

func nthTriangleNumber() {

	for j := 2; ; j++ {
		i := (j * (j + 1)) / 2
		ttlDivisors := 0
		for x := 2; x <= i; x++ {
			if i%x == 0 {
				ttlDivisors++
			}
		}
		if ttlDivisors > 200 {
			println("answer: ", i)
			return
		}
	}
}
