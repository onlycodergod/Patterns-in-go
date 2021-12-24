package pattern

import "fmt"

/*
	Реализовать паттерн "фабричный метод".
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного паттерна на практике
	https://en.wikipedia.org/wiki/Factory_method_pattern


В Go невозможно реализовать классический вариант паттерна Фабричный метод, т.к. в языке отсутствуют возможности ООП, в том числе классы и наследственность.
Но можно реализовать базовую версию этого паттерна — Простая фабрика.

В примере реализации показано, как создавать разные типы оружия при помощи структуры фабрики.
*/

// Интерфейс продукта
// определяет общий интерфейс объектов, которые может произвести создатель и его подклассы.
type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// Конкретные продукты содержат код различных продуктов. Продукты будут отличаться реализацией, но интерфейс у них будет общий.

// Продукт пушки
type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getPower() int {
	return g.power
}

// Конкретный продукт ak47
// включает в себя структуру gun и не напрямую реализуют все методы от iGun
type ak47 struct {
	gun
}

func newAk47() iGun {
	return &ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

// Конкретный продукт - musket
// включает в себя структуру gun и не напрямую реализуют все методы от iGun
type musket struct {
	gun
}

func newMusket() iGun {
	return &musket{
		gun: gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

// Фабрика создает пушку нужного типа в зависимости от аргумента на входе
func gunFactory(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}

// Клиент создает экземпляры различного оружия при помощи gunFactory, используя для контроля изготовления только параметры в виде строк.
func FactoryMethodPatternRun() {
	ak47, _ := gunFactory("ak47")
	musket, _ := gunFactory("musket")
	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
