package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

func removeSpace(input string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, input)
}

func transpose(matrix [][]rune, row int, column int) (oMatrix [][]rune) {
	oMatrix = make([][]rune, column)
	for y, sy := range matrix {
		for x := range sy {
			oMatrix[x] = append(oMatrix[x], matrix[y][x])
		}
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		sentenceWithoutSpace := removeSpace(scanner.Text())
		squareRootOfRuneCount := math.Sqrt(float64(len(sentenceWithoutSpace)))

		var row, column int
		row = int(math.Floor(squareRootOfRuneCount))
		column = int(math.Ceil(squareRootOfRuneCount))

		if row*column < len(sentenceWithoutSpace) {
			row++
		}

		matrix := make([][]rune, row)

		for i := 0; i < row; i++ {
			if i == row-1 {
				matrix[i] = []rune(sentenceWithoutSpace[i*column:])
			} else {
				matrix[i] = []rune(sentenceWithoutSpace[i*column : (i+1)*column])
			}
		}

		oMatrix := transpose(matrix, row, column)

		var strSlice []string

		for i := 0; i < column; i++ {
			strSlice = append(strSlice, string(oMatrix[i]))
		}

		fmt.Println(strings.Join(strSlice[:], " "))
	}
}
