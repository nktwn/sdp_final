package main

import "fmt"

type Strategy interface {
	Execute(data string)
}

type ConcreteStrategyA struct{}

func (s *ConcreteStrategyA) Execute(data string) {
	fmt.Println("ConcreteSrtategyA executed with", data)
}

type ConcreteStrategyB struct{}

func (s *ConcreteStrategyB) Execute(data string) {
	fmt.Println("ConcreteStrategyB executed with", data)
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) DoSomething(data string) {
	c.strategy.Execute(data)
}

func main() {

	context := &Context{}

	strategyA := &ConcreteStrategyA{}
	context.SetStrategy(strategyA)
	context.DoSomething("code for A")

	strategyB := &ConcreteStrategyB{}
	context.SetStrategy(strategyB)
	context.DoSomething("code for B")
}
