package main

import "fmt"

// SubsystemClass представляет класс подсистемы с некоторой сложной логикой.
type SubsystemClass struct{}

func (s *SubsystemClass) operation1() string {
	return "Subsystem 1 operation\n"
}

func (s *SubsystemClass) operation2() string {
	return "Subsystem 2 operation\n"
}

// Facade предоставляет упрощенный интерфейс для сложной системы подсистем.
type Facade struct {
	subsystem1 *SubsystemClass
	subsystem2 *SubsystemClass
}

// NewFacade создает новый фасад с подсистемами.
func NewFacade() *Facade {
	return &Facade{
		subsystem1: &SubsystemClass{},
		subsystem2: &SubsystemClass{},
	}
}

// subsystemOperation обеспечивает упрощенную операцию, которая использует подсистемы.
func (f *Facade) subsystemOperation() {
	fmt.Print(f.subsystem1.operation1())
	fmt.Print(f.subsystem2.operation2())
}

func main() {
	facade := NewFacade()
	facade.subsystemOperation()
}
