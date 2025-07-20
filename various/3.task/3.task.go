package main

import "fmt"
//Задача 2: Что выведет и почему?

func main() {
    funcs := []func(){}

    for i := 0; i < 3; i++ {
        funcs = append(funcs, func() {
            fmt.Println(i)
        })
    }

    for _, f := range funcs {
        f()
    }
}
