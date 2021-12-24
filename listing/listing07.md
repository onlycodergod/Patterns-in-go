Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Вывод: на новых строках сначала числа от 1 до 8, а затем постоянно 0

После того, как горутина внутри asChan() закроет канал, функция merge() будет читать из закрытых каналов. 
Т.к. тип каналов - int, то в бесконечном цикле будут извлекаться нули из закрытых каналов. 
Вывод значений из канала, возвращаемого merge() - в main() в бесконечном цикле с range.
```
