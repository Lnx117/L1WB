package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Первый таймер начал")
	waitFirst(3)
	fmt.Println("Первый таймер закончил")
	fmt.Println("Второй таймер")
	waitSec(3)
	fmt.Println("Второй таймер закончил")
}

//первый вариант
func waitFirst(seconds int) {
	// time.After ожидает истечения продолжительности d, а затем отправляет текущее время по возвращаемому каналу
	timer := time.After(time.Second * time.Duration(seconds))
	//Ждем отправки сообщения (отправит по завершении timer)
	<-timer
}

//второй вариант
func waitSec(seconds int) {
	//создаем таймер
	timer := time.NewTimer(time.Duration(seconds) * time.Second)
	//Ждем отправки сообщения
	<-timer.C
}

/* Разница между time.After и time.NewTimer в том что у второго есть метод Stop который предотвращает срабатывание таймера,
а также в сборе мусора. time.After не восстанавливается сборщиком мусора, пока таймер не истекет, а у time.NewTimer можно вызвать
Stop если таймер больше не нужен */
