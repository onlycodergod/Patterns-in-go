package pattern

import "fmt"

/*
	Реализовать паттерн "строитель".
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного паттерна на практике
	https://en.wikipedia.org/wiki/Builder_pattern


Пример использования на практике: создание конфигураций различных домов
*/

// Интерфейс строителя объявляет шаги конструирования продуктов, общие для всех видов строителей
type iBuilder interface {
	setWindowType()
	setNumFloor()
	getHouse() house
}

func getBuilder(builderType string) iBuilder {
	if builderType == "panel" {
		return &panelBuilder{}
	}

	if builderType == "village" {
		return &villageBuilder{}
	}
	return nil
}

// Конкретный строитель 1. Реализуют строительные шаги, для панельного дома
type panelBuilder struct {
	windowType string
	floor      int
}

func newPanelBuilder() *panelBuilder {
	return &panelBuilder{}
}

func (b *panelBuilder) setWindowType() {
	b.windowType = "plastic"
}

func (b *panelBuilder) setNumFloor() {
	b.floor = 5
}

func (b *panelBuilder) getHouse() house {
	return house{
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// Конкретный строитель 2. Реализуют строительные шаги, для домика в деревне
type villageBuilder struct {
	windowType string
	floor      int
}

func newVillageBuilder() *villageBuilder {
	return &villageBuilder{}
}

func (b *villageBuilder) setWindowType() {
	b.windowType = "wooden"
}

func (b *villageBuilder) setNumFloor() {
	b.floor = 1
}

func (b *villageBuilder) getHouse() house {
	return house{
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// Продукт строительства - создаваемый объект
type house struct {
	windowType string
	floor      int
}

// Директор определяет порядок вызова строительных шагов для производства той или иной конфигурации продуктов
type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

// Процесс построения сервиса конкретным строителем
func (d *director) buildHouse() house {
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

// Клиент передаёт строителя через параметр строительного метода директора
func BuilderPatternRun() {
	panelBuilder := getBuilder("panel")
	villageBuilder := getBuilder("village")

	director := newDirector(panelBuilder)
	panelHouse := director.buildHouse()

	fmt.Printf("Panel House\nwindow: %s\nfloor: %v\n\n", panelHouse.windowType, panelHouse.floor)

	director.setBuilder(villageBuilder)
	villageHouse := director.buildHouse()

	fmt.Printf("Village House.\nwindow: %s\nfloor: %v\n\n", villageHouse.windowType, villageHouse.floor)
}
