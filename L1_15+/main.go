package main

import (
	"fmt"
	"strconv"
)

func main() {
	someFunc()
}

func createHugeString(a int) string {
	return strconv.Itoa(a)
}

func someFunc() {
	var justString string
	v := createHugeString(1 << 10)
	x := len(v)
	justString = v[:x]
	fmt.Println(justString)
}

/* Число v слишком короткое чтобы использовать v[:100], выйдет за диапазон
Нужно уменьшить диапазон или увеличить v

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
} */
