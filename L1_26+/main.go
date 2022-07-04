package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcd"))
	fmt.Println(lengthOfLongestSubstring("abCdefAaf"))
	fmt.Println(lengthOfLongestSubstring("abcCd"))
}

func lengthOfLongestSubstring(s string) bool {
	s = strings.ToLower(s)
	if len(s) < 2 {
		return true
	}

	if len(s) == 2 && s[0] != s[1] {
		return true
	}

	i := 0
	j := 1

	//создаем пустой массив в который будем складывать обозначения наших букв
	//и присваивать им значение 1 если оно уникально
	seen := make(map[byte]int)
	//для первого значения ставим 1 сразу
	seen[s[i]] = 1

	//i и j это наше окно в один шаг
	for j < len(s) && i < len(s) {
		//обращаемся к след элементу, вернет 0 если такого еще нет
		if seen[s[j]] != 0 {
			//если такой есть, то false
			return false
		} else {
			//если мы не нашли такого элемента то он уникальный, значит заносим его со значением 1
			// и инкрементируем правую сторону окна
			seen[s[j]] = 1
			j++
		}
	}

	return true
}
