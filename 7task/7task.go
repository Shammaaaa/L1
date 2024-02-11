package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	data  map[string]interface{}
	mutex sync.RWMutex
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		data: make(map[string]interface{}),
	}
}

func (m *ConcurrentMap) Put(key string, value interface{}) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.data[key] = value
}

func (m *ConcurrentMap) Get(key string) (interface{}, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	value, ok := m.data[key]
	return value, ok
}

func (m *ConcurrentMap) Delete(key string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	delete(m.data, key)
}

func main() {
	cmap := NewConcurrentMap()

	var wg sync.WaitGroup
	wg.Add(2)

	// Запись данных
	go func() {
		defer wg.Done()
		cmap.Put("key1", "value1")
	}()

	go func() {
		defer wg.Done()
		cmap.Put("key2", "value2")
	}()

	wg.Wait()

	// Чтение данных
	value1, ok1 := cmap.Get("key1")
	value2, ok2 := cmap.Get("key2")

	fmt.Println(value1, ok1)
	fmt.Println(value2, ok2)

	// Удаление данных
	cmap.Delete("key1")

	value1, ok1 = cmap.Get("key1")
	fmt.Println(value1, ok1)
}

/*
Создаем структуру ConcurrentMap, которая содержит карту данных data и мьютекс mutex.
Метод Put использует мьютекс для блокировки записи в карту, метод Get использует мьютекс для блокировки чтения из карты,
а метод Delete использует мьютекс для блокировки удаления из карты.
Создаем экземпляр структуры ConcurrentMap с помощью функции NewConcurrentMap().
Затем можно использовать методы Put, Get и Delete для конкурентного доступа к данным карты.
*/
