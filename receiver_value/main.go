package main

import "fmt"

type ValueReceiver struct {
	Data string
}

func (receiver ValueReceiver) SetData(data string) {

	receiver.Data = data
}

func main() {

	asPointer := &ValueReceiver{Data: "Hello"}
	asPointer.SetData("Hello World")
	fmt.Println(asPointer.Data)

	asValue := ValueReceiver{Data: "Hello"}
	asValue.SetData("Hello World")
	fmt.Println(asValue.Data)
}
