package main

import "fmt"

type ProductA interface {
	Use() string
}

type ProductB interface {
	Use() string
}

type ConcreteProductA1 struct{}

func (p *ConcreteProductA1) Use() string {
	return "A1"
}

type ConcreteProductB1 struct{}

func (p *ConcreteProductB1) Use() string {
	return "B1"
}

type Factory interface {
	CreateProductA() ProductA
	CreateProductB() ProductB
}

type ConcreteFactory1 struct{}

func (f *ConcreteFactory1) CreateProductA() ProductA {
	return &ConcreteProductA1{}
}

func (f *ConcreteFactory1) CreateProductB() ProductB {
	return &ConcreteProductB1{}
}

func main() {
	factory := &ConcreteFactory1{}
	productA := factory.CreateProductA()
	productB := factory.CreateProductB()
	fmt.Println(productA.Use())
	fmt.Println(productB.Use())
}
