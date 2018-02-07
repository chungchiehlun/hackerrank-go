package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	bufinput := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscanf(bufinput, "%d", &t)
	for i := 0; i < t; i++ {
		var m, n int
		fmt.Fscanf(bufinput, "\n%d\n%d\n", &m, &n)

		cost := make([]int, n)
		for i := range cost {
			fmt.Fscanf(bufinput, "%d", &cost[i])
		}

		for i, x := range cost {
			for j, y := range cost {
				if j > i && x+y == m {
					fmt.Printf("%d %d\n", i+1, j+1)
				}
			}
		}
	}
}
