package main

import (
	"fmt"
)

func main() {
	var n int32
	n = 4
	fmt.Printf("Value of n : %d\n", n)

	var row int32
	var column int32

	for row = 1; row <= n; row++ {
		for column = 1; column <= n; column++ {
			if column <= n-row {
				fmt.Print(" ")
			}
			if column > n-row {
				fmt.Print("#")
			}
		}
		fmt.Print("\n")
	}
}
