Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
Вывод:
0
1
2
3
4
5
6
7
8
9
fatal error: all goroutines are asleep - deadlock!

Чтение из пустого канала блокирует до записи в этот канал или до его закрытия.
Цикл for i := range c получает значения из канала до тех пор, пока он не закрыт.

Как исправить: 
Закрыть канал close(ch) после цикла внутри горутины.

```
