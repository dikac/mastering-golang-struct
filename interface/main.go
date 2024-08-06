package main

import "fmt"

type Name interface {
	SetName(name string)
	GetName() string
}

type Person struct {
	Name string
}

func (receiver *Person) SetName(name string) {

	receiver.Name = name
}
func (receiver *Person) GetName() string {

	return receiver.Name
}

func main() {

	name := &Person{}
	name.SetName("John")

	Print(name)

}

func Print(name Name) {

	fmt.Println(name.GetName())
}
