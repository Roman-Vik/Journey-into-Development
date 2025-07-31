package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan int)
	ch2 := make(chan int)

	// 2 продюсера
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer close(ch)
		ch <- 100
	}()
	go func() {
		defer wg.Done()
		defer close(ch2)
		ch2 <- 288
	}()

	out := merge(&wg, ch, ch2)

	for v := range out {
		fmt.Println(v)
	}
}

func merge(wg *sync.WaitGroup, cs ...<-chan int) <-chan int {
	out := make(chan int)

	// читаем из всех каналов
	for _, ch := range cs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for data := range ch {
				out <- data
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
