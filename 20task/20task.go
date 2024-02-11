package main

import (
	"fmt"
	"strings"
)

// Функция reverseWords переворачивает слова в строке и возвращает результат
func reverseWords(str string) string {
	words := strings.Fields(str)
	reversedWords := make([]string, len(words))

	// Переворачиваем каждое слово в массиве words
	for i, word := range words {
		reversedWords[len(words)-1-i] = word
	}

	return strings.Join(reversedWords, " ")
}

func main() {
	str := "snow dog sun"
	reversedStr := reverseWords(str)
	fmt.Println(reversedStr)
}

/*
Программа определяет функцию reverseWords, которая принимает строку в качестве аргумента и возвращает строку, в которой слова перевернуты.
Внутри функции мы используем функцию strings.Fields, которая разбивает строку на слова и возвращает массив слов words.
Затем мы создаем новый массив reversedWords той же длины и переворачиваем слова в него, сохраняя порядок слов изначальной строки.
В функции main мы определяем исходную строку str и вызываем функцию reverseWords для получения строки с перевернутыми словами.
Затем мы выводим результат на экран.
*/
