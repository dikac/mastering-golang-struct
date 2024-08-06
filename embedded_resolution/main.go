package main

import "fmt"

func main() {

	registration := Registration{
		Name: "AB 2001",
	}
	registration.Person.SetName("John") // specific mutation
	registration.Car.SetName("Honda")   // specific mutation
	fmt.Println(registration.Person.Name)
	fmt.Println(registration.Car.Name)
	fmt.Println(registration.Name)

	registration.SetName("Audi") // registration mutation
	fmt.Println(registration.Person.Name)
	fmt.Println(registration.Car.Name)
	fmt.Println(registration.Name)
}
