package main

import (
	"fmt"
)

func main() {
	// Заданная последовательность строк
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем мапу для хранения элементов множества
	set := make(map[string]bool)

	// Добавляем каждую строку в множество
	for _, str := range sequence {
		set[str] = true
	}
	keys := make([]string, 0, len(set))
	for i := range set {
		keys = append(keys, i)

	}
	fmt.Println(keys)

}
