// package main

// import "fmt"

// func main() {
// 	var x string = "Hello, World"
// 	fmt.Println(x)
// 	x = "Welcome!"
// 	fmt.Println(x)

// 	const pi float64 = 3.14159
// 	fmt.Println(pi)

// 	z := "Now"
// 	fmt.Println(z)

// 	var (
// 		i = 5
// 		t = 15
// 		a = 10
// 	)
// 	fmt.Println(i, t, a)
// }

package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}
