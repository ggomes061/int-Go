package main

import "fmt"

func main() {
	fmt.Print("Quantos fahrenheit esta agora: ")
	var f float64
	fmt.Scan(&f)

	cel := (f - 32) / 1.8

	fmt.Println("A temperatura e ", cel, "Celcius")
}
