package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	bufinput := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscanf(bufinput, "%d %d\n", &n, &k)

	position := make([]int, n)
	for i := range position {
		fmt.Fscanf(bufinput, "%d", &position[i])
	}

	// sort
	sort.Slice(position, func(i, j int) bool { return position[i] < position[j] })

	coverage := 0
	count := 0
	Pbase := position[0]

	for _, Pi := range position {
		if Pi > coverage {
			Pbase = Pi
			count++
			coverage = Pbase + k
		} else if k >= Pi-Pbase {
			coverage = Pi + k
		}
	}
	fmt.Println(count)
}
