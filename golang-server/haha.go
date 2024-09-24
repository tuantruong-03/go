package main

import (
	"fmt"
	"reflect"
)

func mains() {
	x := struct {
		Foo string
		Bar int
	}{"foo", 2}

	v := reflect.ValueOf(x)
	t := v.Type() // Get the type information

	values := make(map[string]interface{})

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name // Get the field name
		fmt.Print(fieldName)
		fieldValue := v.Field(i).Interface() // Get the field value as an interface{}
		values[fieldName] = fieldValue       // Store the field name and value in the map
	}
}
