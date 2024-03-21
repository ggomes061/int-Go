package main

import "fmt"

func CMedia(v []int) float64 {
	sum := 0
	for _, x := range v {
		sum += x
	}

	return float64(sum) / float64(len(v))
}

func main() {
	v := []int{5, 6, 7, 8, 3}

	x := CMedia(v)
	fmt.Println(x)

}
