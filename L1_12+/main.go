package main

import "fmt"

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(str)
	x := getOwnSet(str)
	fmt.Println(x)

}
func getOwnSet(str []string) []string {
	res := make([]string, 0)
	// Будем использовать key map для хранения уникальных значений множества
	checkMap := make(map[string]int)
	//Одинаковые значения будут перезаписаны, а уникальные созданы
	for _, val := range str {
		checkMap[val] = 1
	}

	//Записываем уникальные значения (ключи мапы) в результирующий массив
	for key := range checkMap {
		res = append(res, key)
	}
	return res
}
