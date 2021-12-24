package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===
Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.
Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}
Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}
start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)
fmt.Printf(“fone after %v”, time.Since(start))
*/

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()

	res := Or(
		sig(time.Second),
		sig(time.Second*2),
		sig(time.Second*3),
		sig(time.Second*5),
	)

	for out := range res {
		fmt.Println(out)
	}
	fmt.Printf("done after %v\n", time.Since(start))
}

// Or объединяет каналы в один
func Or(channels ...<-chan interface{}) <-chan interface{} {
	single := make(chan interface{})
	var wg sync.WaitGroup

	wg.Add(len(channels))
	for _, ch := range channels {
		go func(ch <-chan interface{}) {
			for v := range ch {
				single <- v
			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(single)
	}()

	return single
}
