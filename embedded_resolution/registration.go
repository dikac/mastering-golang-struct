package main

type Registration struct {
	Person
	Car
	Name string
}

// SetName overrides SetName from Person and Car
func (receiver *Registration) SetName(name string) {

	receiver.Person.SetName(name)
	receiver.Car.SetName(name)
	receiver.Name = name
}
