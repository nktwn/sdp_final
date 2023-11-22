package main

import "fmt"

type Component interface {
	Execute() string
}

type ConcreteComponent struct{}

func (c *ConcreteComponent) Execute() string {
	return "ConcreteComponent"
}

type BaseDecorator struct {
	component Component
}

func NewBaseDecorator(c Component) *BaseDecorator {
	return &BaseDecorator{component: c}
}

func (d *BaseDecorator) Execute() string {
	return d.component.Execute()
}

type ConcreteDecorator struct {
	*BaseDecorator
}

func (d *ConcreteDecorator) Execute() string {
	return "<ConcreteDecorator> " + d.BaseDecorator.Execute() + " </ConcreteDecorator>"
}

func main() {
	component := &ConcreteComponent{}

	decorator := NewBaseDecorator(component)
	fmt.Println(decorator.Execute())

	concreteDecorator := &ConcreteDecorator{BaseDecorator: decorator}
	fmt.Println(concreteDecorator.Execute())
}
