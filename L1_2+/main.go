package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := Arr{
		arr: []int{2, 4, 6, 8, 10},
	}

	wg := new(sync.WaitGroup)

	wg.Add(len(arr.arr))

	for i := range arr.arr {
		go GetNumSquare(arr.ReadNum(i), wg)
	}

	wg.Wait()
}

type Arr struct {
	mu  sync.Mutex
	arr []int
}

func (a *Arr) ReadNum(pos int) int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.arr[pos]
}

func GetNumSquare(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	// or math.pow(num, 2)
	numSquare := num * num
	fmt.Println(numSquare)
}
