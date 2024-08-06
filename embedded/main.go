package main

import "fmt"

type Person struct {
	Name string
}

func (receiver *Person) SetName(name string) {

	receiver.Name = name
}

type Employee struct {
	Person // struct Embedding
	Role   string
}

type EmployeePointer struct {
	*Person // struct Embedding as pointer
	Role    string
}

func main() {

	// like any other value, value embedded struct are created with zero value
	valueEmbedded := Employee{}
	valueEmbedded.SetName("Hello Employee Value")
	fmt.Println(valueEmbedded.Name)

	// pointer embedded struct, must be initialized first otherwise will cause panic
	pointerEmbedded := EmployeePointer{Person: &Person{}}
	pointerEmbedded.SetName("Hello Employee Pointer")
	fmt.Println(pointerEmbedded.Name)

	// like any other value, pointer embedded struct are default to nil
	nilPointerEmbedded := EmployeePointer{}
	//nilPointerEmbedded.SetName("Hello World") // panic
	//fmt.Println(nilPointerEmbedded.Name)
	_ = nilPointerEmbedded
}
