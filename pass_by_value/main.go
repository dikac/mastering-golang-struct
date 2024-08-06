package main

import "fmt"

type Value struct {
	Data string
}

func PrintValue(value Value) {
	value.Data = "Hello World Arg"
	fmt.Println(value.Data)
}

func main() {

	var value3 Value

	value1 := Value{Data: "Hello"}
	value2 := value1
	value3 = value1

	value1.Data = "Hello World 1" // mutate the copy of value1 struct
	value2.Data = "Hello World 2" // mutate the copy of value1 struct
	value3.Data = "Hello World 3" // mutate the copy of value1 struct

	PrintValue(value1) // mutate the copy of value1 struct

	fmt.Println(value1.Data)
	fmt.Println(value2.Data)
	fmt.Println(value3.Data)
}
