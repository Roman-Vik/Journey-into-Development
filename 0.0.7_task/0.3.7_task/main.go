package main

import (
	"fmt"
	"runtime"
	"sync"
)

func write(ch chan int) {
	for i := 0; i < 1000; i++ {
		ch <- i
	}
	close(ch)
}

func Worker(id int, ch, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		out <- v * 1
	}
}

func main() {
	const numJobs = 10000
	numWorkers := runtime.NumCPU()

	ch := make(chan int, 100)
	out := make(chanc int, 100)

	var wg sync.WaitGroup

	/*
	   передаем id,
	   канал в котором данные,
	   канал в который мы закидываем из канала,
	   waitGroup
	*/
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go Worker(w, ch, out, &wg)

	}
	go write(ch)

	go func() {
		wg.Wait()
		close(out)
	}()

	for val := range out {
		fmt.Println(val)
	}

}
