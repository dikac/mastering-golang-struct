package main

func main() {

	conflict := Registration{}
	conflict.Name = "Audi"   // compile error
	conflict.SetName("Audi") // compile error

}
