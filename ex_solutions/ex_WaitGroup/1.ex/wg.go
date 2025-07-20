package main

import (
	"sync"
	"context"
	"fmt"
  "time"
	"errors"
)

// Ну окей создали методы для удобного использования
type waiter interface {
  wait() error
  //передаем контекст и функцию с контекстом которя возвращает ошибку
  run(ctx context.Context, f func(ctx context.Context) error) 
}

// wg, канал с ошибки и канал для структуры
type waitGroup struct {
  wg *sync.WaitGroup
  err chan error
  sema chan struct{}
}
// 
func (g *waitGroup) wait() error {  
  go func(){
    g.wg.Wait()
    close(g.err)
  }()
  var e error
  for v := range g.err {
    e = errors.Join(e,v) // конкатенируем ошибки в одну если, а нил отбрасываем
  }
  
  return e
}

func (g *waitGroup) run(ctx context.Context, fn func(ctx context.Context) error) {
  g.sema <- struct{}{} // добавляем в канал пустую структуру
  g.wg.Add(1) 
  go func(){
    defer g.wg.Done()
    select {
      case <-ctx.Done():
        <-g.sema // если контекст завершился, то читаем из канала и завершаем работу
        return
      default:
    } 
    if err := fn(ctx); err != nil {
      g.err <- err
    }
    <-g.sema
  }()
}

func newGroupWait(maxParallel int) waiter {
  return &waitGroup{
    wg: &sync.WaitGroup{},
    err: make(chan error),
    sema: make(chan struct{},maxParallel),
  }
}
func main() {
  time.Sleep(time.Second*10)
  g := newGroupWait(2) // имплементируем структуру который реализуем waiter
  ctx := context.Background()
  
  // создали 2 ошибки
  expErr1 := errors.New("got error 1") 
  expErr2 := errors.New("got error 2")
  

  g.run(ctx, func(ctx context.Context) error {
    return nil
  })
  g.run(ctx, func(ctx context.Context) error {
    return expErr2
  })
  g.run(ctx, func(ctx context.Context) error {
    return expErr1
  })
  err := g.wait()
  if !errors.Is(err, expErr1) || !errors.Is(err, expErr2) {
    panic("wrong code")
  } else {
    fmt.Println("ok")
  }
}