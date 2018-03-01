package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	bufin := bufio.NewReader(os.Stdin)
	var p, d, m, s int
	fmt.Fscanf(bufin, "%d %d %d %d\n", &p, &d, &m, &s)

	output := 0
	totalCost := 0

	for totalCost < s {
		output++
		price := p - d*(output-1)

		if price < m {
			price = m
		}

		totalCost += price

		if totalCost > s {
			output = output - 1
		}
	}

	fmt.Println(output)
}
