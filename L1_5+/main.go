package main

import (
	"fmt"
	"time"
)

func main() {
	var secBeforeExit int
	fmt.Print("Время работы программы: ")
	fmt.Scan(&secBeforeExit)

	ch := make(chan string)

	go WriteToChan(ch)
	go ReadFromChan(ch)
	//Запускаем горутины и ждем нужное время
	time.Sleep(time.Duration(secBeforeExit) * time.Second)
}

//Пишем в канал с небольшой задержкой для наглядности
func WriteToChan(ch chan string) {
	for {
		ch <- "ping"
		time.Sleep(250 * time.Millisecond)
	}
}

//Считываем с канала и выводим данные
func ReadFromChan(ch chan string) {
	for {
		data := <-ch
		fmt.Println(data)
	}
}
