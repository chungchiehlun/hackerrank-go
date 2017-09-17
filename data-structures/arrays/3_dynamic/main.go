package main

import "fmt"

func main() {
	var lenthOfSeqList, numOfQueries, lastAnswer, queryType, x, y int
	fmt.Scan(&lenthOfSeqList)
	fmt.Scan(&numOfQueries)
	seqList := make([][]int, lenthOfSeqList)
	for i := 1; i <= numOfQueries; i++ {
		fmt.Scan(&queryType)
		fmt.Scan(&x)
		fmt.Scan(&y)
		index := (x ^ lastAnswer) % lenthOfSeqList
		switch queryType {
		case 1:
			seqList[index] = append(seqList[index], y)
		case 2:
			lastAnswer = seqList[index][y%len(seqList[index])]
			fmt.Printf("%d\n", lastAnswer)
		}
	}
}
