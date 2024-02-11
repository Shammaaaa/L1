package main

import "fmt"

// Определение структуры Human
type Human struct {
	name string
	age  int
}

func (h Human) Info() {
	fmt.Printf("Name: %s\nAge: %d\n", h.name, h.age)
}

// Определение структуры Action
type Action struct {
	Human
}

func main() {
	// Создание экземпляра структуры Action
	action := Action{
		Human: Human{
			name: "Shamil",
			age:  24,
		},
	}

	// Вызов метода
	action.Info()

}

/*
Структура Action встраивает структуру Human с использованием встроенного полиморфизма.
Это означает, что структура Action получает все поля и методы структуры Human
*/
