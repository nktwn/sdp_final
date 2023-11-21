package main

import (
	"fmt"
)

// Component интерфейс, который определяет методы, которые должны реализовать и декорируемый объект, и декоратор.
type Component interface {
	Operation() string
}

// ConcreteComponent конкретная реализация компонента, которую мы будем декорировать.
type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
	return "ConcreteComponent"
}

// Decorator базовая структура декоратора, встраивает Component для дальнейшего расширения.
type Decorator struct {
	component Component
}

func (d *Decorator) Operation() string {
	return d.component.Operation()
}

// ConcreteDecoratorA конкретный декоратор, который добавляет новое поведение к Component.
type ConcreteDecoratorA struct {
	Decorator
}

func (d *ConcreteDecoratorA) Operation() string {
	return "<ConcreteDecoratorA>" + d.component.Operation() + "</ConcreteDecoratorA>"
}

// ConcreteDecoratorB еще один конкретный декоратор, который добавляет свое поведение.
type ConcreteDecoratorB struct {
	Decorator
}

func (d *ConcreteDecoratorB) Operation() string {
	return "<ConcreteDecoratorB>" + d.component.Operation() + "</ConcreteDecoratorB>"
}

func main() {
	// Создаем конкретный компонент
	var component Component = &ConcreteComponent{}

	// Декорируем компонент декоратором A
	decoratorA := &ConcreteDecoratorA{Decorator{component}}
	fmt.Println(decoratorA.Operation())

	// Декорируем компонент декоратором B
	decoratorB := &ConcreteDecoratorB{Decorator{component}}
	fmt.Println(decoratorB.Operation())

	// Можно декорировать декоратор, создавая сложные структуры
	decoratedComponent := &ConcreteDecoratorA{
		Decorator: Decorator{
			component: &ConcreteDecoratorB{
				Decorator: Decorator{
					component: component,
				},
			},
		},
	}
	fmt.Println(decoratedComponent.Operation())
}
