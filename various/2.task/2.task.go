package main

import "fmt"

type account struct{
	value int
}
//Задача 1: Что выведет и почему?

func main() {
	// 
    s1 := make([]account, 0, 2) 
	
    s1 = append(s1, account{})
	
    s2 := append(s1, account{})
fmt.Println(s1, s2)
    acc := &s2[0]
    acc.value = 10
   // fmt.Println(s1, s2)
     acc.value += 10
     //fmt.Println(s1, s2)
}

