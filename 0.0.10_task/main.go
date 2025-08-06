package main

import (
	"fmt"
	"math/rand"
	"time"
)

func uniqRandom(n int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // инициализируем генератор
	var res []int

	for i := 0; i < n; i++ {
		res = append(res, r.Intn(100))
	}

	return res
}

func main() {
	fmt.Println(uniqRandom(10))
}
