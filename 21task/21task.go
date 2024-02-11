package main

import "fmt"

// Интерфейс "Target"
type Target interface {
	Request()
}

// Интерфейс "Adaptee"
type Adaptee interface {
	SpecificRequest()
}

// Структура, реализующая интерфейс "Adaptee"
type ConcreteAdaptee struct{}

func (ca *ConcreteAdaptee) SpecificRequest() {
	fmt.Println("Вызван метод SpecificRequest")
}

// Структура "Adapter", реализующая интерфейс "Target"
type Adapter struct {
	adaptee Adaptee
}

func (a *Adapter) Request() {
	a.adaptee.SpecificRequest()
}

// Функция для использования "Target"
func UseTarget(target Target) {
	target.Request()
}

func main() {
	adaptee := &ConcreteAdaptee{}
	adapter := &Adapter{adaptee}
	UseTarget(adapter)
}

/*
При выполнении программы будет вызван метод Request() через адаптер, который запускает метод SpecificRequest() интерфейса Adaptee.
В результате будет выведено сообщение "Вызван метод SpecificRequest".
*/
