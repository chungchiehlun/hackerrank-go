package main

import "fmt"

func main() {
	var N, Q int
	fmt.Scan(&N)
	slice := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&slice[i])
	}
	fmt.Scan(&Q)
	for i := 0; i < Q; i++ {
		var snippet string
		fmt.Scan(&snippet)
		fmt.Println(search(slice, snippet))
	}
}

func search(slice []string, snippet string) int {
	count := 0
	for _, value := range slice {
		if value == snippet {
			count++
		}
	}
	return count
}
