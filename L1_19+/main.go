package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "123abcабв"

	//рассплитили стринг
	splitedStr := strings.Split(input, "")

	n := len(splitedStr)

	//идем к середине меняя левый с правым
	for i := 0; i < n/2; i++ {
		splitedStr[i], splitedStr[n-1-i] = splitedStr[n-1-i], splitedStr[i]
	}

	fmt.Println(splitedStr)
}
