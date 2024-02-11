package main

import (
	"fmt"
)

// Функция для чтения чисел из канала и умножения их на 2
func multiply(input chan int) chan int {
	output := make(chan int)

	go func() {
		for num := range input {
			output <- num * 2
		}
		close(output)
	}()

	return output
}

func main() {
	// Исходный массив чисел
	numbers := []int{1, 2, 3, 4, 5}

	// Канал для записи чисел из массива
	input := make(chan int)

	// Запускаем функцию для умножения чисел на 2
	doubled := multiply(input)

	// Записываем числа из массива в канал input
	go func() {
		for _, num := range numbers {
			input <- num
		}
		close(input)
	}()

	// Выводим числа из второго канала
	for num := range doubled {
		fmt.Println(num)
	}
}

/*
Функция multiply создает новый канал output и запускает горутину, которая читает числа из канала input,
умножает их на 2 и записывает в канал output.
Функция main создает канал input, запускает горутину для записи чисел из массива в канал input,
а затем читает числа из канала doubled и выводит их.
*/
