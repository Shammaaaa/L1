package main

import (
	"fmt"
	"reflect"
)

func main() {
	var variable interface{}

	// Примеры различных типов переменных
	variable = 42 // int
	fmt.Println("Тип переменной:", getType(variable))

	variable = "Привет" // string
	fmt.Println("Тип переменной:", getType(variable))

	variable = true // bool
	fmt.Println("Тип переменной:", getType(variable))

	c := make(chan int)
	variable = c // channel
	fmt.Println("Тип переменной:", getType(variable))
}

func getType(v interface{}) string {

	return reflect.TypeOf(v).String()
}
