package main

import "fmt"

func main() {

	t := []float64{20, 12, 18, 28}
	var x []float64
	y := make([]float64, 5)
	z := make([]float64, 8, 10)

	arr := [5]float64{1, 2, 3, 1, 5}
	//x := arr[0:5]

	slice := []int{1, 2, 3}
	slice1 := append(slice, 5, 7)
	slice2 := make([]int, 5)

	q := make(map[string]int)
	q["key"] = 10
	fmt.Println(q["key"])

	w := make(map[int]int)
	w[3] = 10
	fmt.Println(w[3])

	fmt.Println(t, x, y, z, arr, slice1, slice2)
}
