package main

import ("fmt";"sync")

func worker(chIn <- chan int, chOut chan <- int, done chan <- struct{}){
	var wg sync.WaitGroup
	
	for v := range chIn{
		wg.Add(1)
		go func(val int){
			defer wg.Done()
			chOut <- val * 100
		}(v)
	}

	go func(){
		wg.Wait()
		close(chOut)
		done <- struct{}{}
	}()
}

func main(){
	chIn := make(chan int, 3)
	chOut := make(chan int, 3)
	done := make(chan struct{})
	
// Отправляем данные в chIn
	chIn <- 1
	chIn <- 30
	close(chIn)

	go worker(chIn, chOut, done)

	for v := range chOut{
		fmt.Println(v)
	}
	<- done
	fmt.Println("Done!")
}