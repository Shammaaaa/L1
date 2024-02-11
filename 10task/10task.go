package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// Заданные значения температуры
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создаем словарь для хранения групп температур
	groups := make(map[float64][]float64)

	// Сортируем значения температуры по возрастанию
	sort.Float64s(temperatures)

	// Обрабатываем каждое значение температуры
	for _, temp := range temperatures {
		// Определяем ключ группы для текущей температуры
		key := math.Floor(temp/10) * 10

		// Добавляем текущую температуру в соответствующую группу
		groups[key] = append(groups[key], temp)
	}

	// Выводим группы температур
	for key, group := range groups {
		fmt.Printf("Группа %.0f градусов:\n", key)
		fmt.Println(group)
	}
}
