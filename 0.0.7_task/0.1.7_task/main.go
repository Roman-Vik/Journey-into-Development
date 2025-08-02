package main

import "fmt"

func worker(done chan struct{}){
	fmt.Println("Работник: жду сигнал для завершения...")
	<- done //! Ждём сигнал из канала
	fmt.Println("Работник: 'получил сигнал, завершаю работу.", )
}

func main(){
	done := make(chan struct{},3)

	go worker(done)

	fmt.Println("отправляю сигнал в канал")
	done <-struct{}{}
	 fmt.Println("Главная: программа завершена")
}