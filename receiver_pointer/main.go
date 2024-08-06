package main

import "fmt"

type PointerReceiver struct {
	Data string
}

func (receiver *PointerReceiver) SetData(data string) {

	receiver.Data = data
}

func main() {

	asPointer := &PointerReceiver{Data: "Hello"}
	asPointer.SetData("Hello World")
	fmt.Println(asPointer.Data)

	asValue := PointerReceiver{Data: "Hello"}
	asValue.SetData("Hello World")
	fmt.Println(asValue.Data)

}
