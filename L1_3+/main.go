package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	arr := Array{
		arr: [5]int{2, 4, 6, 8, 10},
	}

	wg := new(sync.WaitGroup)

	//Ждем горутины считывающие массив и одну читающую
	wg.Add(len(arr.arr) + 1)

	for i := range arr.arr {
		go arr.GetNumSquare(i, ch, wg)
	}

	//Читает канал пока тот не закроется, суммирует все что приходит, после этого выводит
	go ReadChan(ch, wg)

	wg.Wait()
}

type Array struct {
	mu  sync.Mutex
	arr [5]int
}

//Функция конкурентного считывания
func (a *Array) ReadNum(pos int) int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.arr[pos]
}

//Конкурентно считывает со слайса, а затем возводит в квадрат и скидывает полученное в канал
func (a *Array) GetNumSquare(pos int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	num := a.ReadNum(pos)
	// or math.pow(num, 2)
	numSquare := num * num
	ch <- numSquare

	//Закрываем канал на последнем элементе массива
	if pos == len(a.arr)-1 {
		close(ch)
	}
}

//Читает канал. Суммирует пока канал открыт
func ReadChan(ch chan int, wg *sync.WaitGroup) {

	var sum int
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println(sum)
			break
		} else {
			sum = sum + val
		}
	}
	defer wg.Done()
}
