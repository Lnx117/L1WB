package main

import (
	"fmt"
	"reflect"
)

func main() {
	var intVar int
	var stringVar string
	var floatVar float32
	var b bool
	fmt.Printf("Переменная типа - %s\n", varType(intVar))
	fmt.Printf("Переменная типа - %s\n", varType(stringVar))
	fmt.Printf("Переменная типа - %s\n", varType(floatVar))
	fmt.Printf("Переменная типа - %s\n", varType(b))
}

func varType(variable interface{}) reflect.Type {
	return reflect.TypeOf(variable)
}
