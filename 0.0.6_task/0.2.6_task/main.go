package main

//Задача 1: Написать код функции, которая делает merge N каналов. Весь входной поток перенаправляется в один канал.
//Без гпт и других источников! Задача базовая
import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// Write

	ch := make(chan int)
	ch2 := make(chan int)
	go func() {
		defer close(ch)
		ch <- 100
	}()
	go func() {
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

	for _, ch := range cs {
		wg.Add(1)
		// обязательно передаем ch как аргумент, иначе замыкание!
		go func(ch <-chan int) {
			defer wg.Done()
			for data := range ch {
				out <- data
			}
		}(ch)
	}
	// закрытие канала out после завершения всех горутин
	go func() {
		wg.Wait()
		close(out)

	}()

	return out
	// Write
}
