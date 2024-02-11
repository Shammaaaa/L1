package main

import (
	"fmt"
	"sync"
)

// Структура Counter представляет счетчик со счетчиком и мьютексом для обеспечения конкурентного доступа
type Counter struct {
	value int
	mutex sync.Mutex
}

// Метод Increment инкрементирует счетчик
func (c *Counter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.value++
}

// Метод GetValue возвращает текущее значение счетчика
func (c *Counter) GetValue() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.value
}

func main() {
	counter := Counter{} // Создание экземпляра счетчика

	// Создание группы ожидания для ожидания завершения всех горутин
	var wg sync.WaitGroup

	wg.Add(3) // Установка количества горутин, которые нужно дождаться

	// Запуск горутин, которые инкрементируют счетчик
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			counter.Increment()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 101; i++ {
			counter.Increment()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 60; i++ {
			counter.Increment()
		}
	}()

	wg.Wait() // Ожидание завершения всех горутин
	fmt.Println("Итоговое значение счетчика:", counter.GetValue())
}

/*
Приведенный код создает структуру Counter, которая содержит счетчик и мьютекс для синхронизации доступа к счетчику.
Метод Increment инкрементирует счетчик, а метод GetValue возвращает текущее значение счетчика.
В функции main создается экземпляр Counter и запускается несколько горутин, каждая из которых инкрементирует счетчик какое-то кол-во раз.
Затем программа ожидает завершения всех горутин и выводит итоговое значение счетчика.
*/
