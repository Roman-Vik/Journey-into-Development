package main

import "fmt"

//Задача 2: Что выведет и почему?

func main() {
	// слайс функций
    funcs := []func(){}
	// аппендим слайс функций
    for i := 0; i < 3; i++ {
        funcs = append(funcs, func() {
            fmt.Println(i)
        })
    }
	// вызываем слайс функций 
    for _, f := range funcs {
        f()
    }
}
// вывод: 0,1,2