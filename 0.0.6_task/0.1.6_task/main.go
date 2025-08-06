package main

import (
	"fmt"
	"sync"
)

/*
Задача 1: Написать код функции, которая делает merge N каналов. Весь входной поток перенаправляется в один канал.
Без гпт и других источников! Задача базовая
*/
func main() {

	ch := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch)
		ch <- 1
		ch <- 10
	}()

	go func() {
		defer close(ch2)
		ch2 <- 2
	}()

	out := merge(ch, ch2)

	for val := range out {
		fmt.Println(val)
	}

	// Write
}

func merge(cs ...<-chan int) <-chan int {

	var wg sync.WaitGroup
	out := make(chan int)

	for _, ch := range cs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for val := range ch {
				out <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
	// Write
}
