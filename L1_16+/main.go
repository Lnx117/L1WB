package main

import (
	"fmt"
)

func main() {
	fmt.Println(quickSort([]int{1, 3, 7, 2, 9, 1, 12, 9}, 0, 7))
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	//Опорный элемент самый правый
	pivot := arr[high]
	//Точка делящая массив на более маленькие. Также является левой частью плавающего окна
	i := low
	//Проходим по массиву слева направо
	for j := low; j < high; j++ {
		//Если опорный элемент больше текущего то меняем текущий и меньший элементы местами
		//Таким образом мы сдвигаем элементы больше опорного к правой части, а те что меньше к левой
		//Наткнувшись на элемент больше опорного i будет ждать пока не найдется элемент меньше и потом поменяет местами больший с меньшим
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	//В последней смене i указывает на элемент больше чем опорный
	arr[i], arr[high] = arr[high], arr[i]
	//Возвращаем новый массив и индекс опорного элемента
	return arr, i
}
