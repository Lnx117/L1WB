package main

import "fmt"

func main() {
	arr := []int{1, 2, 5, 6, 7, 9, 12, 17, 19, 44, 65, 78, 98}
	fmt.Println(search(arr, 17))
}

func search(nums []int, target int) int {
	right := len(nums) - 1 //индекс последнего эл-та в массиве
	left := 0              // индекс первого

	for left <= right { //пока левый индекс меньше либо равен правом
		pivot := (left + right) / 2 //находим середину отрезка
		if target == nums[pivot] {  // если в середине отрезка то что мы ищем то конец
			return pivot
		} else if target > nums[pivot] { //если то что мы ищем больше чем то что в середине, то сдвигаем левый край к след от этого элемента (потом будет новая середина и тд)
			left = pivot + 1
		} else {
			right = pivot - 1 //если нет то сдвигаем правый край
		}
	}
	return -1
}
