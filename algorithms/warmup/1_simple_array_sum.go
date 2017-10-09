package main

import "fmt"

func readStdin(length int) []int {
	s := make([]int, length)
	for i := range s {
		_, err := fmt.Scan(&s[i])
		if err != nil {
			fmt.Println(err)
		}
	}
	return s
}

func main() {
	var length int
	fmt.Scan(&length)
	slice := readStdin(length)
	sum := slice[0]
	for i := 1; i < length; i++ {
		sum = sum + slice[i]
	}
	fmt.Println(sum)
}
