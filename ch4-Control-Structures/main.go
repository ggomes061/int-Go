package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}

	for o := 1; o <= 10; o++ {
		fmt.Println(o)
		if o == 5 {
			fmt.Println("WARNIG!")
			break
		} else if o > 5 {
			fmt.Println("Acabou")
		}
	}

	for z := 1; z <= 10; z++ {
		switch z {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Five")
		case 5:
			fmt.Println("Four")
		}
	}

	for r := 3; r <= 100; r += 3 {
		fmt.Println(r)
	}

}
