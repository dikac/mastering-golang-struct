package main

import (
	"fmt"
	"reflect"
)

type WithTag struct {
	Data string `custom:"data"`
}

func main() {

	reflection := reflect.TypeOf(WithTag{})
	reflectionField, _ := reflection.FieldByName("Data")
	tag := reflectionField.Tag.Get("custom")
	fmt.Println(tag)

}
