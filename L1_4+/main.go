package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var workerCount int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&workerCount)

	//Для signal.Notify требуется буферизированный канал, 1 будет достаточно
	/* 	Package signal will not block sending to c: the caller must ensure that c has sufficient
	buffer space to keep up with the expected signal rate.  */
	exitChan := make(chan os.Signal, 1)
	//При нажатии ctrl + c (SIGINT) в канал sigCh будет отправлено сообщение
	signal.Notify(exitChan, syscall.SIGINT)
	go Exit(exitChan)

	//Для бесконечной работы воркеров
	wg := new(sync.WaitGroup)
	wg.Add(1)

	ch := make(chan int)
	for i := 0; i <= workerCount; i++ {
		go Worker(ch, i)
	}

	go WriteToChan(ch)

	wg.Wait()
}

//Публикуем данные в канал (просто числа)
func WriteToChan(ch chan int) {
	dataCounter := 0
	for {
		ch <- dataCounter
		dataCounter++
		time.Sleep(100 * time.Millisecond)
	}
}

//Читаем данные из канала и выводим
func Worker(ch chan int, num int) {
	for {
		data := <-ch
		fmt.Println("Воркер номер: ", num, "Считал с канала: ", data)
	}
}

//Ждем нужного сигнала для выхода из программы
func Exit(exitChan chan os.Signal) {
	for {
		switch <-exitChan {
		case syscall.SIGINT:
			fmt.Println("User press ctrl + C .... exit")
			os.Exit(0)
		default:
		}
	}
}
