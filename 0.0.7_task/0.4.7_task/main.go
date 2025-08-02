package main

import (
	"fmt"
	"runtime"
)

func worker(read, change chan int, done chan struct{}) {
	for val := range read {
		change <- val
	}

	done <- struct{}{}
}

func main() {
	task := 100

	ch, change, done := make(chan int), make(chan int), make(chan struct{})

	for i := 1; i <= runtime.NumCPU(); i++ {
		go worker(ch, change, done)
	}

	go func() {
		for job := 1; job <= task; job++ {
			ch <- job
		}
		close(ch)
	}()

	// Ожидаем завершения всех воркеров и закрываем `change`
	go func() {
		for i := 1; i <= runtime.NumCPU(); i++ {
			<-done
		}
		close(change)
	}()

	
	for res := range change {
		fmt.Println(res)
	}
	

}
