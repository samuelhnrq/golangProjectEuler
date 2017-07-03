package main

func prob15() {
	side := 20 + 1
	grid := make([][]int, side)
	for i := range grid {
		grid[i] = make([]int, side)
		grid[i][0] = 1
		grid[0][i] = 1
	}
	for i := 1; i < side; i++ {
		for x := 1; x < side; x++ {
			grid[x][i] = grid[x-1][i] + grid[x][i-1]
		}
	}
	println(grid[side-1][side-1])
}
