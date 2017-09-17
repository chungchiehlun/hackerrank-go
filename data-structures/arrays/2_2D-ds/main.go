package main

import "fmt"

func main() {
	fmt.Println(
		maximum(
			calculateHourglassValue(
				readStdin(),
			),
		),
	)
}

func readStdin() (array [6][6]int) {
	for x := 0; x < 6; x++ {
		for y := 0; y < 6; y++ {
			fmt.Scan(&array[x][y])
		}
	}
	return
}

func calculateHourglassValue(array [6][6]int) (hourglass [16]int) {
	for i := 0; i < 16; i++ {
		type C struct {
			x int
			y int
		}
		var center C
		center.x = i/4 + 1
		center.y = i%4 + 1
		// fmt.Printf("x:%d, y:%d\n", center.x, center.y)
		hourglass[i] = array[center.x-1][center.y-1] +
			array[center.x-1][center.y] +
			array[center.x-1][center.y+1] +
			array[center.x][center.y] +
			array[center.x+1][center.y-1] +
			array[center.x+1][center.y] +
			array[center.x+1][center.y+1]
	}
	return
}

func maximum(values [16]int) (max int) {
	max = values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}
