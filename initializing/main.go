package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

	{ // defining as value will also initialize struct to zero value
		var value User
		fmt.Println(value.Name)
		fmt.Println(value.Age)
	}

	{ // initialize with shorthand assignment
		value := User{}
		fmt.Println(value.Name)
		fmt.Println(value.Age)
	}

	{ // defining as pointer will initialize to nil
		var value *User
		_ = value
		//fmt.Println(value.Name)
		//fmt.Println(value.Age)
	}

	{
		// initialize as pointer
		pointer := &User{}
		fmt.Println(pointer.Name)
		fmt.Println(pointer.Age)
	}

	{ // initialize with value
		value := User{
			Name: "John Doe",
			Age:  21,
		}
		fmt.Println(value.Name)
		fmt.Println(value.Age)
	}

	{
		// initialize as pointer with value
		pointer := &User{
			Name: "Johnny Doe",
			Age:  22,
		}
		fmt.Println(pointer.Name)
		fmt.Println(pointer.Age)
	}
}
