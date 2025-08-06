package main

import (
	"context"
	"log"
	"runtime/debug"
	"time"
)

func SafeGoCtx(ctx context.Context, fn func(ctx context.Context)) {
	go func() {
		defer func() {
			if r := recover(); r != nil {     //debug.Stack() выводит стек вызовов.
				log.Printf("panic: %v\n%s", r, debug.Stack())
			}
		}()
		fn(ctx)
	}()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go SafeGoCtx(ctx, func(ctx context.Context) {
		panic("dcmcd")
	})

	time.Sleep(time.Second)
}
