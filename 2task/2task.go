package main

import (
	"fmt"
	"sync"
)

func Square(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := num * num
	fmt.Println(result)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go Square(num, &wg)
	}

	wg.Wait()
}

/*
Мы определяем функцию Square, которая рассчитывает квадрат числа и выводит результат. Она принимает число и указатель на объект sync.WaitGroup.
В функции main мы создаем массив чисел [2, 4, 6, 8, 10] и итерируемся по нему.
Для каждого числа мы вызываем функцию Square в отдельной горутине, передавая число и указатель на объект sync.WaitGroup.
После запуска всех горутин мы вызываем метод Wait у объекта sync.WaitGroup, чтобы дождаться их завершения.
*/
