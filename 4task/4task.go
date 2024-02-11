package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func worker(id int, dataChan <-chan string, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case data := <-dataChan:
			fmt.Printf("%d получил: %s\n", id, data)
		case <-ctx.Done():
			fmt.Printf("%d закончил\n", id)
			return
		}
	}
}

func main() {
	numWorkers := 3                             // Количество воркеров
	data := []string{"data1", "data2", "data3"} // Произвольные данные

	var wg sync.WaitGroup
	dataChan := make(chan string)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Запуск воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChan, &wg, ctx)
	}

	// Запись данных в канал
	for _, d := range data {
		dataChan <- d
	}

	// Перехват сигналов прерывания
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Отмена контекста и ожидание завершения всех воркеров
	cancel()
	wg.Wait()
	fmt.Println("Все закончили")
}

/*
В этом примере мы создаем корневой контекст с помощью context.Background().
Затем мы вызываем функцию WithCancel, чтобы создать контекст ctx и функцию cancel, которую мы будем вызывать для прерывания работы программы.
В функции воркера, вместо использования for и select для чтения данных из канала и ожидания отмены контекста,
мы используем select для выбора между чтением данных из канала и получением сигнала об отмене контекста.
Далее, мы создаем канал sigChan для перехвата сигналов прерывания (Ctrl+C) и используем signal.Notify,
чтобы отправить сигналы os.Interrupt, syscall.SIGINT и syscall.SIGTERM в канал sigChan.
Затем мы ожидаем получения сигнала из канала sigChan.
При получении сигнала о прерывании мы вызываем функцию cancel для отмены контекста,
что приводит к завершению работы всех воркеров, ожидающих этого контекста.
Затем мы ожидаем завершения работы всех воркеров с помощью wg.Wait()
*/
