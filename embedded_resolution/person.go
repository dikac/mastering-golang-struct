package main

type Person struct {
	Name string
}

func (receiver *Person) SetName(name string) {

	receiver.Name = name
}
