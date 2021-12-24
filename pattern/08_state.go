package pattern

import (
	"fmt"
	"log"
)

/*
	Реализовать паттерн "состояние".
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного паттерна на практике
	https://en.wikipedia.org/wiki/State_pattern



// Контекст
// хранит ссылку на объект состояния и делегирует ему часть работы, зависящей от состояний.
// Контекст работает с этим объектом через общий интерфейс состояний. Контекст должен иметь метод для присваивания ему нового объекта-состояния.
type vendingMachine struct {
	hasItem       state
	itemRequested state
	noItem        state
	currentState  state // состояние
	itemCount     int
}

func newVendingMachine(itemCount int) *vendingMachine {
	v := &vendingMachine{
		itemCount: itemCount,
	}
	hasItemState := &hasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &itemRequestedState{
		vendingMachine: v,
	}
	noItemState := &noItemState{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.noItem = noItemState
	return v
}

func (v *vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *vendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

func (v *vendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}

// Интерфейс состояния описывает общий интерфейс для всех конкретных состояний.
type state interface {
	addItem(int) error
	requestItem() error
	dispenseItem() error
}

// Конкретный интерфейс состояния "Нет предмета"
type noItemState struct {
	vendingMachine *vendingMachine
}

func (i *noItemState) requestItem() error {
	return fmt.Errorf("item out of stock")
}

func (i *noItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

func (i *noItemState) dispenseItem() error {
	return fmt.Errorf("item out of stock")
}

// Конкретный интерфейс "Имеет предмет"
type hasItemState struct {
	vendingMachine *vendingMachine
}

func (i *hasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("no item present")
	}
	fmt.Printf("Item requestd\n")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

func (i *hasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
}

func (i *hasItemState) dispenseItem() error {
	return fmt.Errorf("please select item first")
}

// Конкретный интерфейс "Выдать предмет"
type itemRequestedState struct {
	vendingMachine *vendingMachine
}

func (i *itemRequestedState) requestItem() error {
	return fmt.Errorf("item already requested")
}

func (i *itemRequestedState) addItem(count int) error {
	return fmt.Errorf("item Dispense in progress")
}

func (i *itemRequestedState) dispenseItem() error {
	fmt.Println("Dispensing Item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}

// Клиент
func StatePatternRun() {
	vendingMachine := newVendingMachine(1)

	err := vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.addItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
