package main

import (
    "fmt"
    "time"
)

// worker получает сигнал через канал-сигнализатор и завершает работу
func worker(done chan struct{}) {
    fmt.Println("Работник: жду сигнала для завершения...")
    <-done // Ждём сигнал из канала
    fmt.Println("Работник: получил сигнал, завершаю работу.")
}

func main() {
    done := make(chan struct{}) // Канал для сигналов, ничего не содержит

    go worker(done) // Запуск горутины-работника

    time.Sleep(2 * time.Second)
    fmt.Println("Главная: отправляю сигнал для завершения")
    done <- struct{}{} // Отправляем сигнал

    time.Sleep(1 * time.Second) // Даем горутине завершиться
    fmt.Println("Главная: программа завершена")
}
