package main

import (
	"fmt"
)

//Задача 2: Как сделать чтобы первый case выбирался всегда?

func main() {
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)

	close(ch1)
	close(ch2)
	ch1Value := 0.0
	ch2Value := 0.0

	for i := 0; i < 100000; i++ {
		select {
		case <-ch1:
			ch1Value++
		default:
			select {
			case <-ch2:
				ch2Value++
      default:
			}
		}
	}

	fmt.Println(ch1Value / ch2Value) // ожидается 100000 / 0
}
/*
Да, потому что:
    ch1Value == 100000
    ch2Value == 0
    float64(100000) / 0 == +Inf
*/