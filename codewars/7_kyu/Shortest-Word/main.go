package main

import (
	"fmt"
	"regexp"
)

func FindShort(s string) int {
	re, err := regexp.Compile(`\s+`)
	if err != nil {
		fmt.Println("Ошибка компиляции:", err)
	}
	data := re.Split(s, -1)

	count := data[0]

	for _, v := range data {
		if len(v) < len(count) {
			count = v
		}
	}

	return len(count)
}

func main() {
	res := FindShort("Hello Wor1ld Roma")
	fmt.Println(res)
}
