package pattern

import "fmt"

/*
	Реализовать паттерн "команда".
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного паттерна на практике
	https://en.wikipedia.org/wiki/Command_pattern

Пример использования: реализация кнопок ВКЛ/ВЫКЛ телевизора с помощью пульта и кнопкой на экране.
Сначала можно реализовать команду ВКЛ, где телевизор является Получателем. Когда на эту команду вызывается метод execute, она вызывает функцию TV.on.
Реализацию этого определяет вызывающий объект, которых будет два: пульт и телевизор. Оба будут содержать объект команды ВКЛ.
Мы обернули запрос в несколько вызывающих объектов. Преимущество создания отдельных объектов команд в отделении логики пользовательского интерфейса от внутренней бизнес-логики.
Нет нужды создавать отдельные исполнители для каждого вызывающего объекта – сама команда содержит всю информацию, необходимую для ее исполнения.
Соответственно, ее можно использовать для отсроченного выполнения задачи.
*/

// Отправитель
// Хранит ссылку на объект команды и обращается к нему, когда нужно выполнить какое-то действие.
// Отправитель работает с командами только через их общий интерфейс. Он не знает, какую конкретно команду использует, т.к. получает готовый объект команды от клиента.
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

// Интерфейс команды
// Описывает общий для всех конкретных команд интерфейс. Обычно здесь описан всего один метод для запуска команды.
type command interface {
	execute()
}

// Конкретные команды
// Реализуют различные запросы, следуя общему интерфейсу команд.
// Обычно команда не делает всю работу самостоятельно, а лишь передаёт вызов получателю, которым является один из объектов бизнес-логики.

// Конкретная команда ВКЛ
type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

// Конкретная команда ВЫКЛ
type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

// Интерфейс получателя
// Содержит бизнес-логику программы
type device interface {
	on()
	off()
}

// Конкретный получатель - TV
type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

// Клиентский код
// создаёт объекты конкретных команд, передавая в них все необходимые параметры, среди которых могут быть и ссылки на объекты получателей.
// После этого клиент связывает объекты отправителей с созданными командами.
func CommandPatternRun() {
	tv := &tv{}
	onCommand := &onCommand{device: tv}
	offCommand := &offCommand{device: tv}

	onButton := &button{command: onCommand}
	onButton.press()

	offButton := &button{command: offCommand}
	offButton.press()
}
