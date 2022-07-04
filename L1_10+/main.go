package main

import "fmt"

func main() {
	input := []float64{-25.4, -27.6, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	fmt.Println(MakeGroups(input))
}

func MakeGroups(arr []float64) map[int][]float64 {
	TempMap := make(map[int][]float64)
	for _, v := range arr {
		//Таким образом мы берем только десятичную часть числа и получаем ключ по которому определяем к какой группе отнести температуру
		//Затем записываем числов в эту группу
		num := int(v) / 10 * 10
		TempMap[num] = append(TempMap[num], v)
	}
	return TempMap
}
