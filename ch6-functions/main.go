package main

import "fmt"

func zero(zZ *int) {
	*zZ = 0
}

func one(xOne *int) {
	*xOne = 1
}

func main() {
	x := 8
	zero(&x)
	fmt.Println(x)

	xZ := new(int)
	one(xZ)
	fmt.Println(*xZ)
}
