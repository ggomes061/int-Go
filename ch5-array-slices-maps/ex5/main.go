package main

import "fmt"

func main() {
	// x := [6]string{"a", "b", "c", "d", "e", "f"}

	// fmt.Println(x[2:5]) //cde

	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	m := finMin(x)
	fmt.Println("O menor numero e: ", m)
}

func finMin(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	min := arr[0]

	for _, num := range arr {
		if num < min {
			min = num
		}
	}

	return min
}
