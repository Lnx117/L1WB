package main

import "fmt"

func main() {
	x := 10
	newX := setBit(x, 0)
	fmt.Println(newX)
	newX = clearBit(x, 0)
	fmt.Println(newX)
}

/* Создаем маску с битом равным 1 на том месте где нужно его установить, затем используем OR (дает ноль только в случае когда
оба бита нули), получим что изменится только необходимый нам бит */
func setBit(n int, pos uint) int {
	n |= (1 << pos)
	return n
}

/* Создаем маску с битом равным 1 на том месте где нужно его установить и инвертируем каждый бит в маске (0010 will be 1101),
затем используем AND (даст единицу только в случае когда оба бита равны единице), получим что изменится только необходимый нам бит */
func clearBit(n int, pos uint) int {
	mask := ^(1 << pos)
	n &= mask
	return n
}
