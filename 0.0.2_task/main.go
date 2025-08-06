package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

type waiter interface {
	wait() error
	run(ctx context.Context, f func(ctx context.Context) error)
}

type waitGroup struct {
	wg   *sync.WaitGroup
	err  chan error
	sema chan struct{}
}

// Этот метод читает ошибки из канала и объединяет в одну и возвращает её
func (g *waitGroup) wait() error {
	go func() {
		g.wg.Wait()  // Ждем что-то?
		close(g.err) // как дождались закрываем канал
	}()
	var e error
	for v := range g.err { // читаем из канала ошибку
		e = errors.Join(e, v) // конкатенируем ошибку
	}
	return e
}

// Метод напоминает воркер
func (g *waitGroup) run(ctx context.Context, fn func(ctx context.Context) error) {
	g.sema <- struct{}{} // передаем в сему структуру
	g.wg.Add(1)          //
	go func() {
		defer g.wg.Done()

		select {
		case <-ctx.Done():
			<-g.sema // читаем из семы если контекст завешился
			return   // значит ошибки объеденены
		default: //! мне кажется лишний он тут не нужен
		}
		if err := fn(ctx); err != nil {
			g.err <- err // пишем в канал ошибку
		}
		<-g.sema // чтитаем из семы
	}()
}

func newGroupWait(maxParallel int) waiter { // имплементируем методы
	return &waitGroup{
		wg:   &sync.WaitGroup{},
		err:  make(chan error),
		sema: make(chan struct{}, maxParallel),
	}
}

func main() {
	maxParallel := 2
	g := newGroupWait(maxParallel)
	// контекст
	ctx := context.Background() //! Зачем контекст если его не используем?
	// errors.New  возвращает ошибку
	expErr1 := errors.New("got error 1")
	expErr2 := errors.New("got error 2")

	g.run(ctx, func(ctx context.Context) error {
		return nil // нет ошибки сработает паника
	})
	g.run(ctx, func(ctx context.Context) error {
		return expErr1
	})
	g.run(ctx, func(ctx context.Context) error {
		return expErr2
	})

	err := g.wait() // тут возвращаются ошибки
	//  если хотя бы одна ошибка не найдена  сработает паника
	if !errors.Is(err, expErr1) || !errors.Is(err, expErr2) {
		panic("wronng code")
	} else {
		fmt.Println("ok")
	}
}
