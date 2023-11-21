package main

import "fmt"

// Product определяет интерфейс продукта, который фабрика будет создавать.
type Product interface {
	Use() string
}

// ConcreteProductA конкретная реализация продукта A.
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string {
	return "Product A is used"
}

// ConcreteProductB конкретная реализация продукта B.
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string {
	return "Product B is used"
}

// Factory определяет интерфейс фабрики с методом CreateProduct.
type Factory interface {
	CreateProduct(productType string) Product
}

// ConcreteFactory конкретная реализация фабрики.
type ConcreteFactory struct{}

func (f *ConcreteFactory) CreateProduct(productType string) Product {
	switch productType {
	case "A":
		return &ConcreteProductA{}
	case "B":
		return &ConcreteProductB{}
	default:
		return nil
	}
}

func main() {
	factory := &ConcreteFactory{}

	// Создание продукта типа A
	productA := factory.CreateProduct("A")
	fmt.Println(productA.Use())

	// Создание продукта типа B
	productB := factory.CreateProduct("B")
	fmt.Println(productB.Use())
}
