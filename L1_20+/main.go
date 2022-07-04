package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	example := "door sun phone"
	fmt.Println(changeWordOrder(example))
}
func changeWordOrder(str string) string {
	//Создаем маску для разбитие пословно
	//\s - пробел, + повторение того что в скобках >1 раза, [] любой символ из тех что в скобках
	regExp := regexp.MustCompile(`[\s]+`)

	//-1 значит что сплитим всю строку
	stringArr := regExp.Split(str, -1)

	var res string

	//записываем в строку идя с конца массива  к началу
	for i := len(stringArr) - 1; i >= 0; i-- {
		res += stringArr[i] + " "
	}

	return strings.TrimSuffix(res, " ")
}
