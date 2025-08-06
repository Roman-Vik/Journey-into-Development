package main

import (
	"time"
	"log"
	"runtime/debug"
)

func SafeGo(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {    //debug.Stack() выводит стек вызовов.
				log.Printf("panic: %v\n%s", r, debug.Stack())
			}
		}()
		fn()
	}()
}

func main() {
	SafeGo(func() {
		panic("Boom!")
	})

	time.Sleep(500 * time.Millisecond)
}
