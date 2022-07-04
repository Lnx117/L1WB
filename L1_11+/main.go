package main

import (
	"fmt"
)

func main() {

	ArrX := []int{3, 5, 1, 32, 65, 9, 16, 2, 1, 1, 1}
	ArrY := []int{4, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 1, 1, 1, 1, 1}

	fmt.Println(FindSameElements(ArrX, ArrY))
}

func FindSameElements(ArrX, ArrY []int) []int {
	counter := make(map[int]int)

	//Проходим по первому массиву и записываем в ключ значение элемента, получим список уникальных элементов
	for _, val := range ArrX {
		counter[val] = 1
	}

	var res []int

	for _, val := range ArrY {
		//Проходим по второму массиву и смотрим были ли такие значения в ключах первого. Если такой ключ есть
		// и он еще не был использован (значение равно 1) то пишем в результирующий массив и обнуляем чтобы больше такие числа не записывать
		//
		if xVal, ok := counter[val]; ok && xVal != 0 {
			res = append(res, val)
			counter[val] = 0
		}
	}

	return res
}
