package listing

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Task 01
func Task1() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b) // [77, 78, 79]
}

// Task 02
func Test() (x int) { // returned variable x
	defer func() {
		x++ // x=2
	}()
	x = 1
	return // 2
}

func AnotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x // 1
}

// Task 03
func Foo() error {
	var err *os.PathError = nil
	return err
}

func Task3() {
	err := Foo()
	fmt.Println(err)        // <nil>
	fmt.Println(err == nil) // false
	fmt.Println(err.(*os.PathError) == nil)
}

// Task 04
func Task4() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		//close(ch)
	}()

	for n := range ch {
		println(n)
	}
}

// Task 05
type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

// Вариант исправления 1
// Вместо возврата *customError, возвращаем обычный интерфейс error
// func test() error {
// 	{
// 		// do something
// 	}
// 	return nil
// }

func Task5() {
	var err error
	err = test()

	// Вариант исправления 2
	// интерфейс содержащий нулевой указатель не равен nil, но если возвращать структуру то равен
	//err2 := test()
	//if err2 != nil {

	if err != nil {
		println("error ")
		return
	}
	println("ok")
}

// Task 06
func Task6() {
	var s = []string{"1", "2", "3"}
	modifySlice(s) // сделали одно изменение, а затем потеряли из-за вместимости изначальный слайс и работали с другим
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"         // s = [3, 2, 3]
	i = append(i, "4") // не хватает вместимости в базовом слайсе, поэтому создан новый и в него копируется предыдущий i = [3, 2, 3, 4]
	i[1] = "5"         // с прошлого шага работаем с новым слайсом и изменения делаем в нём i = [3, 5, 3, 4]
	i = append(i, "6") // i = [3, 5, 3, 4, 6]
	//fmt.Println("i", i) // можно посмотреть слайс, с которым мы делали все изменения (кроме первого)
}

// Task 07
// передача чисел по небуферизованному каналу
func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c) // закрытие канала
	}()
	return c
}

// передача чисел из 2-х каналов в 3-й
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

func Task7() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
