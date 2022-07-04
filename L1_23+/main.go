package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c", "d"}

	//первый вариант
	fmt.Println(letters)
	fmt.Println(DeleteElement(3, letters))

	//второй вариант
	fmt.Println(letters)
	fmt.Println(DeleteElementSecondVar(3, letters))

}

//Первый вариант. Просто получаем новый массив с элементами до удаляемого значения и после
func DeleteElement(index int, arr []string) []string {
	if index >= len(arr) || index < 0 {
		fmt.Println("Slice is too small or index < 0")
		return arr
	}
	//... это распаковка  arr[index+1:] так как элементы должны быть переданы поотдельности
	arr = append(arr[0:index], arr[index+1:]...)
	return arr
}

//Второй вариант удаления, заменяет все начиная с удаляемого элемента на то что идет после него
//при этом получится дубликат последнего элемента, удаляем его
func DeleteElementSecondVar(index int, arr []string) []string {
	copy(arr[index:], arr[index+1:])
	arr = arr[:len(arr)-1]
	return arr
}
