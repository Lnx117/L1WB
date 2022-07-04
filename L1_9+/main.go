package main

import (
	"fmt"
)

func main() {
	//канал с числами
	naturals := make(chan int)
	//канал с квадратами
	squares := make(chan int)

	arr := []int{1, 2, 3, 4, 5}

	//анонимная горутина отсылает все данные в обычный канал
	go func() {
		for x := range arr {
			naturals <- arr[x]
		}
		close(naturals)
	}()

	//анонимная горутина принимает данные с канала, удваивает их и отправляет в канал с квадратами чисел
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	//читаем канал с квадратами чисел и выводим их
	for x := range squares {
		fmt.Println(x)
	}
}
