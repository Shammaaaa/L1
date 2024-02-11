package main

import (
	"fmt"
	"sync"
)

func Square(num int, wg *sync.WaitGroup, squareChan chan int) {
	defer wg.Done()
	result := num * num
	squareChan <- result
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	squareChan := make(chan int)

	for _, num := range numbers {
		wg.Add(1)
		go Square(num, &wg, squareChan)
	}

	go func() {
		wg.Wait()
		close(squareChan)
	}()

	sum := 0
	for square := range squareChan {
		sum += square

	}

	fmt.Println(sum)
}

/*
В этом примере мы определяем функцию Square, которая рассчитывает квадрат числа и отправляет результат в канал squareChan.
Также принимает указатель на объект sync.WaitGroup.
В функции main мы создаем массив чисел [2, 4, 6, 8, 10] и итерируемся по нему.
Для каждого числа мы вызываем функцию Square в отдельной горутине, передавая число, указатель на объект sync.WaitGroup и канал squareChan.
Мы также запускаем анонимную горутину, которая ожидает завершения всех горутин и закрывает канал squareChan.
После этого мы пробегаемся по каналу squareChan с помощью цикла for range, суммируя все полученные квадраты чисел.
*/
