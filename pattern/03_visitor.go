package pattern

import "fmt"

/*
	Реализовать паттерн "посетитель".
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного паттерна на практике
	https://en.wikipedia.org/wiki/Visitor_pattern



Пример использования на практике: добавление для структуры фигур круга и квадрата функции вычисления площади и нахождения средней точки
*/

// Элемент описывает метод принятия посетителя. Этот метод должен иметь единственный параметр, объявленный с типом интерфейса посетителя.
type shape interface {
	getType() string
	accept(visitor)
}

// Конеретные элементы Круга и Квадрата
// Реализуют методы принятия посетителя. Цель метода — вызвать тот метод посещения, который соответствует типу этого элемента.
// Так посетитель узнает, с каким именно элементом он работает.

// Конкретный элемент - Круг.
type circle struct {
	radius int
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

func (c *circle) getType() string {
	return "Circle"
}

// Конкретный элемент - Квадрат
type square struct {
	side int
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

func (s *square) getType() string {
	return "Square"
}

// Посетитель описывает общий интерфейс для всех типов посетителей.
// Он объявляет набор методов, отличающихся типом входящего параметра, которые нужны для запуска операции для всех типов конкретных элементов.
type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
}

// Конкретные посетители - нахождение площади и средней точки фигур.
// Реализуют особенное поведение для всех типов элементов, которые можно подать через методы интерфейса посетителя.

// Конкретный посетитель - нахождение площади фигуры
type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	fmt.Println("Calculating area for square")
}

func (a *areaCalculator) visitForCircle(s *circle) {
	fmt.Println("Calculating area for circle")
}

// Конкретный посетитель - нахождение средней точки фигуры
type middleCoordinates struct {
	x int
	y int
}

func (a *middleCoordinates) visitForSquare(s *square) {
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *middleCoordinates) visitForCircle(c *circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}

// Клиентский код
func VisitorPatternRun() {
	circle := &circle{radius: 3}
	square := &square{side: 2}

	areaCalculator := &areaCalculator{}
	circle.accept(areaCalculator)
	square.accept(areaCalculator)
	fmt.Println()

	middleCoordinates := &middleCoordinates{}
	circle.accept(middleCoordinates)
	square.accept(middleCoordinates)
}
