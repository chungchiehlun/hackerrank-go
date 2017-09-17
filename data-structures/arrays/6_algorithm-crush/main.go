package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	list := make([]int, n)
	for i := 0; i < m; i++ {
		var a, b, k int
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&k)
		list[a-1] += k
		if b < n {
			list[b] -= k
		}
	}
	// fmt.Println(strings.Trim(fmt.Sprintf("%v", list), "[]"))
	MAX := 0
	sum := 0
	for i := 0; i < n; i++ {
		sum += list[i]
		if MAX < sum {
			MAX = sum
		}
	}
	fmt.Println(MAX)
}
