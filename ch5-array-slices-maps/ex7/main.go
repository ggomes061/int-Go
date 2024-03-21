package main

import "fmt"

func main() {
	nArray := [4]int{1, 2, 3, 4}
	sArray := [3]string{"Hello", "World", "Main"}
	nArray[0] = 50

	fmt.Println(nArray)
	fmt.Println(sArray)

	for x := range nArray {
		x++

		fmt.Println(x)
	}

}
