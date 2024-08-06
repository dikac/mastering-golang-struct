package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

	// as value
	value := User{}
	value.Name = "John Dean"
	value.Age = 23
	fmt.Println(value.Name)
	fmt.Println(value.Age)

	// as pointer
	pointer := &User{}
	pointer.Name = "John Depp"
	pointer.Age = 24
	fmt.Println(pointer.Name)
	fmt.Println(pointer.Age)

}
