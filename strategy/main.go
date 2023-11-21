package main

import "fmt"

// Strategy интерфейс определяет способ выполнения алгоритма.
type Strategy interface {
	Execute(data string)
}

// ConcreteStrategyA реализация стратегии A.
type ConcreteStrategyA struct{}

func (s *ConcreteStrategyA) Execute(data string) {
	fmt.Println("ConcreteStrategyA executed with", data)
}

// ConcreteStrategyB реализация стратегии B.
type ConcreteStrategyB struct{}

func (s *ConcreteStrategyB) Execute(data string) {
	fmt.Println("ConcreteStrategyB executed with", data)
}

// Context содержит ссылку на стратегию.
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) DoSomething(data string) {
	c.strategy.Execute(data)
}

// main функция демонстрирует использование паттерна Strategy.
func main() {
	context := Context{}

	strategyA := &ConcreteStrategyA{}
	context.SetStrategy(strategyA)
	context.DoSomething("data for A")

	strategyB := &ConcreteStrategyB{}
	context.SetStrategy(strategyB)
	context.DoSomething("data for B")
}
