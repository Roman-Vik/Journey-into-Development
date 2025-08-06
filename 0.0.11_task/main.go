package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Задача:
Необходимо написать параллельную обработку в кол-ве горутин равном workers.
Получаем данные из канала in,
над каждым элементом выполняем processData.
Передаем результат в out
В дополнение к этому необходимо добавить получение контекста и его обработку
*/

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	in := make(chan int, 100)
	out := make(chan int, 100)

	// Generating input data
	go func() {
		for i := 0; i < 100; i++ {
			in <- i
		}
		close(in)
	}()

	processInParallel(ctx, in, out, 5)
	var count int
	for result := range out {
		count++
		fmt.Println("Processed:", result, count)
	}
}

func processInParallel(ctx context.Context, in <-chan int, out chan<- int, workers int) {
	wg := sync.WaitGroup{}
	// Write here
	for job := 0; job < workers; job++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, in, out)
		}()
	}

	go func() {
		wg.Wait()
		close(out) // Закрытие out после завершения всех воркеров через sync.WaitGroup.
	}()
}

func worker(ctx context.Context, in <-chan int, out chan<- int) {
	for {
		select {
		//  если отменился контект, то выходим
		case <-ctx.Done():
			return
			// проверяем значения в канале
		case val, ok := <-in:
			// закончились значения выходим из программы
			if !ok {
				return
			}
			select {
			//  если отменился контект, то выходим
			case <-ctx.Done():
				return
			// записываем в нанал измененные значения
			case out <- processData(val):
			}
		}
	}
}

func processData(data int) int {
	time.Sleep(100 * time.Millisecond) // Simulate time-consuming task
	return data * 2
}
