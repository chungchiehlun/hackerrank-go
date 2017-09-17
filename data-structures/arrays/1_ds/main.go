package main

import "fmt"

func main() {
	var length int
	fmt.Scan(&length)
	slice := readStdin(length)
	reverseSlice(slice)
	printSlice(slice)
}

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

// reverse order of slice
func reverseSlice(a []int) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

// Printing slice without the '[' ']'
func printSlice(a []int) {
	for i, v := range a {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
