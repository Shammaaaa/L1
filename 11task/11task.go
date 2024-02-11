package main

import "fmt"

func main() {
	// Первое множество
	set1 := []int{1, 3, 5, 7, 9}

	// Второе множество
	set2 := []int{2, 3, 5, 7, 8}

	// Создаем мапу для хранения элементов и их наличия
	elements := make(map[int]bool)

	// Добавляем элементы первого множества в мапу
	for _, element := range set1 {
		elements[element] = true
	}

	// Создаем слайс для хранения пересечения элементов
	intersection := []int{}

	// Проверяем каждый элемент второго множества на наличие в мапе
	for _, element := range set2 {
		if elements[element] {
			intersection = append(intersection, element)
		}
	}

	fmt.Println("Пересечение множеств:", intersection)
}
