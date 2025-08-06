package main

import "log"

//Задача 3: Что выведет и почему?

type MyStruct struct{}
//метод у труктуры который выводит дату со временем и 
func (s MyStruct) Hello() {
    log.Println("Hello")
}

func main() {
	// Инициализируем переменную a
    a := MyStruct{}
	// Инициализируем переменную b и передаем указатель
    b := &a
	// и у b по сути это a и у него есть вызывам метод
    b.Hello()
}