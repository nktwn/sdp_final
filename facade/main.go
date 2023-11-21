package main

import "fmt"

// Подсистема 1
type SubsystemOne struct{}

func (s *SubsystemOne) OperationOne() string {
	return "SubsystemOne: Готов!\n"
}

func (s *SubsystemOne) OperationN() string {
	return "SubsystemOne: Действие!\n"
}

// Подсистема 2
type SubsystemTwo struct{}

func (s *SubsystemTwo) OperationOne() string {
	return "SubsystemTwo: Готов к запуску!\n"
}

func (s *SubsystemTwo) OperationZ() string {
	return "SubsystemTwo: Огонь!\n"
}

// Фасад
type Facade struct {
	one *SubsystemOne
	two *SubsystemTwo
}

func NewFacade() *Facade {
	return &Facade{
		one: &SubsystemOne{},
		two: &SubsystemTwo{},
	}
}

func (f *Facade) Operation() {
	fmt.Print("Фасад инициализирует подсистемы:\n")
	fmt.Print(f.one.OperationOne())
	fmt.Print(f.two.OperationOne())
	fmt.Print("Фасад заказывает подсистемам выполнение действий:\n")
	fmt.Print(f.one.OperationN())
	fmt.Print(f.two.OperationZ())
}

func main() {
	facade := NewFacade()
	facade.Operation()
}
