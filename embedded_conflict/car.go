package main

type Car struct {
	Name string
}

func (receiver *Car) SetName(name string) {

	receiver.Name = name
}
