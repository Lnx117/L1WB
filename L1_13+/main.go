package main

import (
	"fmt"
)

func main() {
	Changer()
}

func Changer() {
	x, y := 11, 99
	fmt.Printf("x = %d, y = %d \n", x, y)
	x, y = y, x
	fmt.Printf("x = %d, y = %d \n", x, y)
}
