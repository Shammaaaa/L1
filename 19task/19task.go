package main

import (
	"fmt"
	"unicode/utf8"
)

// Функция reverseString переворачивает строку и возвращает результат
func reverseString(str string) string {
	runes := []rune(str)
	length := utf8.RuneCountInString(str)

	// Переворачиваем символы в массиве runes
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	str := "главрыба"
	reversedStr := reverseString(str)
	fmt.Println(reversedStr)
}

/*
Программа определяет функцию reverseString, которая принимает строку в качестве аргумента и возвращает перевернутую версию этой строки.
Для обращения к символам строки, представленным в виде Unicode, мы используем пакет unicode/utf8.
В функции main мы определяем исходную строку str и вызываем функцию reverseString для получения перевернутой версии.
Затем мы выводим перевернутую строку на экран.
*/
