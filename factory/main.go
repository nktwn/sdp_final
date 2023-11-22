package main

import "fmt"

type Product interface {
	DoStuff() string
}

type ConcreteProductA struct{}

func (p *ConcreteProductA) DoStuff() string {
	return "ConcreteProductA stuff"
}

type ConcreteProductB struct{}

func (p *ConcreteProductB) DoStuff() string {
	return "ConcreteProductB stuff"
}

type Creator interface {
	CreateProduct() Product
}

type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) CreateProduct() Product {
	return &ConcreteProductA{}
}

type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func someOperation(creator Creator) {
	product := creator.CreateProduct()
	fmt.Println(product.DoStuff())
}

func main() {
	creatorA := &ConcreteCreatorA{}
	someOperation(creatorA)
	
	creatorB := &ConcreteCreatorB{}
	someOperation(creatorB)
}
