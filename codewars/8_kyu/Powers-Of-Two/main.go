package main

import (
	"fmt"
)

func PowersOfTwo(n int) []uint64 {
	x := []uint64{}
	//
	for i := 0; i <= n; i++ {
		if i >= 3{
			fmt.Println(i+1)
			if i % 2 == 0{
				i+= 2
				fmt.Println(i+2)
			}
		}

	}
	return x
}

func main() {
	PowersOfTwo(4)	
}
