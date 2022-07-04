package main

import (
	"fmt"
	"sync"
)

func main() {
	cnt := &Counter{
		num: 0,
	}

	wg := new(sync.WaitGroup)
	Do(cnt, wg)

	wg.Wait()
	fmt.Printf("Законили работу со значением %d", cnt.num)

}

//счетчик и мьютекс
type Counter struct {
	num int
	sync.Mutex
}

//блочим структуру и инкрементим, затем разблокировываем
func (c *Counter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.num++
}

//геттер
func (c *Counter) Value() int {
	return c.num
}

//запускаем 20 анонимных горутин, каждая инкрементит счетчик когда может
func Do(cnt *Counter, wg *sync.WaitGroup) {

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(num int, cnt *Counter, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Printf("Горутина номер %d начала работу\n", num)
			cnt.Increment()
			fmt.Printf("Горутина номер %d закончила работу\n", num)
		}(i, cnt, wg)
	}
}
