package main

import "fmt"

func main() {
	action := Action{
		SomeData: "data",
		Human: Human{
			Name: "Mike",
			Age:  36,
			Sex:  "Male",
		},
	}

	// Action own method
	action.SayData()
	// Action inherit methods
	action.Human.SayName()
	action.Human.SayAge()

}

type Human struct {
	Name string
	Age  int
	Sex  string
}

func (h *Human) SayName() {
	fmt.Printf("My name is %s \n", h.Name)
}

func (h *Human) SayAge() {
	fmt.Printf("My age is %d \n", h.Age)
}

type Action struct {
	SomeData string
	Human    Human
}

func (a *Action) SayData() {
	fmt.Printf("Data is %s \n", a.SomeData)
}
