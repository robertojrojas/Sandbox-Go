package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	FirstName string `default:"John"`
	LastName  string `default:"Doe"`
	Age       int `default:"30"`
}

func (someone *Person) indentifyYourSelf() {
	val := reflect.ValueOf(someone).Elem()

	for i := 0; i < val.NumField(); i++ {

		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", 
			               typeField.Name, 
			               valueField.Interface(), 
			               tag.Get("default"))
	}
}

func main () {
	roberto := &Person{
		FirstName: "Roberto",
		LastName: "Rojas",
		Age: 44,
	}
    roberto.indentifyYourSelf()

	johnDoe := &Person{}
    johnDoe.indentifyYourSelf()

}