package main

import (
	"fmt"
	"math/big"
)

func main() {
	x, y := new(big.Int), new(big.Int)
	x.SetString("600000000000000000000000000000000", 10)
	y.SetString("200000000000000000000000000000000", 10)

	fmt.Println("X + Y = ", Add(x, y))
	fmt.Println("X - Y = ", Sub(x, y))
	fmt.Println("X * Y = ", Mul(x, y))
	fmt.Println("X / Y = ", Div(x, y))

}

//Просто используем методы big

func Add(x *big.Int, y *big.Int) *big.Int {
	res := new(big.Int)
	res.Add(x, y)
	return res
}
func Sub(x *big.Int, y *big.Int) *big.Int {
	res := new(big.Int)
	res.Sub(x, y)
	return res
}
func Mul(x *big.Int, y *big.Int) *big.Int {
	res := new(big.Int)
	res.Mul(x, y)
	return res
}
func Div(x *big.Int, y *big.Int) *big.Int {
	res := new(big.Int)
	res.Div(x, y)
	return res
}
