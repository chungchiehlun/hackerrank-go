package main

import "fmt"

func main() {
	var n, d int
	fmt.Scan(&n)
	fmt.Scan(&d)
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}
	for i, v := range append(slice[d:], slice[:d]...) {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
