package main

import "fmt"

func main() {
	m := generateMatrix(5)
	fmt.Println(m)
	for _, row := range m {
		for _, e := range row {
			fmt.Printf("%d\t", e)
		}
		fmt.Println()
	}
}

func generateMatrix(n int) [][]int {
	ret := make([][]int, n, n)
	for running := 1; running <= n*n; running++ {
		i, j := next(n, running)
		if len(ret[i]) == 0 {
			ret[i] = make([]int, n, n)
		}
		ret[i][j] = running
	}
	return ret
}

func next(n, running int) (i, j int) {
	running--
	i = running / n
	j = running % n
	fmt.Printf("running = %d, i = %d, j = %d\n", running+1, i, j)
	return
}
