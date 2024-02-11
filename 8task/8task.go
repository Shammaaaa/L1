package main

import (
	"fmt"
)

func main() {
	var num int64
	var i int
	var bitValue int

	// Ввод значения переменной и номера бита
	fmt.Print("Введите значение")
	fmt.Scan(&num)

	fmt.Print("Введите номер бита")
	fmt.Scan(&i)

	fmt.Print("Введите значение бита")
	fmt.Scan(&bitValue)

	// Установка значения бита
	num = setBit(num, i, bitValue)

	// Вывод результата
	fmt.Println(num)
}

// Функция для установки значения i-го бита в num
func setBit(num int64, i int, bitValue int) int64 {
	// Проверка допустимости значения i
	if i < 0 || i > 63 {
		fmt.Println("Недопустимое значение")
		return num
	}

	// Проверка допустимости значения bitValue
	if bitValue != 0 && bitValue != 1 {
		fmt.Println("Недопустимое значение")
		return num
	}

	// Установка значения бита
	if bitValue == 1 {
		num = num | (1 << i) // Установка i-го бита в 1
	} else {
		num = num &^ (1 << i) // Установка i-го бита в 0
	}

	return num
}

/*

 */
