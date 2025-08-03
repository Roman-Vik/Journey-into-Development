package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	w := 5 // сколько запускаем кол-во воркеров

	dataArr := getData() // слайс строк в виде чисел

	chanJob := make(chan string, len(dataArr))

	//наполняем наш буф канал данными из ArrData
	for i := 0; i < len(dataArr); i++ {
		chanJob <- dataArr[i]
	}
	// после того как заполнили канал закрываем
	close(chanJob)

	// Далее используем синхронизацию waitgroup
	wg := sync.WaitGroup{}
	//нам надо создать в цикле и запустить 5 воркеров и передать в них канал chanJob

	for i := 0; i < w; i++ {
		wg.Add(1)
		go func() {
			// после как воркер отработает сработает  wg.Done()
			defer wg.Done() //
			worker(chanJob)
		}()
	}
	wg.Wait()
}

func getData() []string {
	return []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"}
}

func checkDomain(host string) bool {
	if len(host) > 0 {
		return true
	}
	return false
}

// ворекеры запускаются пачками по 5 шт с интервалом в одну секунду
func worker(job chan string) {
	for v := range job {
		bo := checkDomain(v)        // возвращаем тру
		fmt.Println(bo)             // булевое значение вернет функция
		time.Sleep(time.Second *) // задержка в одну секунду
	}
}
