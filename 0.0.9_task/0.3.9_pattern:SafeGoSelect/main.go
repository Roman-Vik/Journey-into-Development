package main

import (
	"context"
	"log"
	"runtime/debug"
	"time"
)

// SafeGoCtx запускает функцию в горутине и защищает от panic
func SafeGoCtx(ctx context.Context, fn func(ctx context.Context)) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[panic recovered] %v\n%s", r, debug.Stack())
			}
		}()
		fn(ctx)
	}()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("[main] starting worker")

	// Запускаем безопасную фоновую горутину
	SafeGoCtx(ctx, func(ctx context.Context) {
		tick := time.NewTicker(500 * time.Millisecond)
		defer tick.Stop()

		for {
			select {
			case <-ctx.Done():
				log.Println("[worker] shutdown gracefully")
				return
			case t := <-tick.C:
				log.Println("[worker] tick at", t.Format(time.StampMilli))

				// Имитируем случайную панику
				if t.Second()%5 == 0 {
					panic("simulated crash in worker")
				}
			}
		}
	})

	// Дадим воркеру поработать 7 секунд
	time.Sleep(7 * time.Second)

	log.Println("[main] cancelling context")
	cancel()

	// Ждём, чтобы воркер успел завершиться
	time.Sleep(1 * time.Second)

	log.Println("[main] done")
}
