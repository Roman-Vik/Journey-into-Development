package main

import "fmt"

//! panic: assignment to entry in nil map

func MapIsNil() {
	var m map[string]int
	fmt.Println(m == nil)

	m["one"] = 10
	fmt.Println(len(m)) //

	fmt.Println(m)
}

func main() {
	MapIsNil()
}
