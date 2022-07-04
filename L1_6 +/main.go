package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	stopByChan(wg)
	stopByFlag(wg)

	wg.Wait()
}

func stopByChan(wg *sync.WaitGroup) {
	ch := make(chan int)
	//Работает пока в канал не будет отправлено сообщение
	go firstWorker(ch, wg)
	time.Sleep(time.Second * 2)
	//Отправляем сообщение и останавливаем первую горутину
	ch <- 1
	close(ch)
}

// первая горутина, остнавливаем каналом
func firstWorker(ch chan int, wg *sync.WaitGroup) {
	for {
		select {
		case <-ch:
			fmt.Println("ПЕРВАЯ ГОРУТИНА ОСТАНОВЛЕНА")
			wg.Done()
			return
		default:
			fmt.Println("Первая горутина что-то делает")
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func stopByFlag(wg *sync.WaitGroup) {
	flag := false
	go secondWorker(&flag, wg)
	time.Sleep(time.Second * 2)
	flag = true
}

// вторая горутина, останавливаем флагом. передаем указатель на переменную и следим за ее состоянием
func secondWorker(stopFlag *bool, wg *sync.WaitGroup) {
	for {
		if *stopFlag {
			fmt.Println("ВТОРАЯ ГОРУТИНА ОСТАНОВЛЕНА")
			wg.Done()
			return
		} else {
			fmt.Println("Вторая горутина что-то делает")
			time.Sleep(250 * time.Millisecond)
		}
	}
}
