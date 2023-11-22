package main

import "fmt"

type Command interface {
	Execute()
}

type Receiver struct{}

// Operation некоторая операция, которую нужно выполнить.
func (r *Receiver) Operation() {
	fmt.Println("Receiver: выполнение операции.")
}

// ConcreteCommand конкретная реализация команды с ссылкой на Receiver.
type ConcreteCommand struct {
	receiver *Receiver
}

// NewConcreteCommand конструктор для ConcreteCommand.
func NewConcreteCommand(receiver *Receiver) *ConcreteCommand {
	return &ConcreteCommand{receiver: receiver}
}

// Execute вызов операции Receiver.
func (c *ConcreteCommand) Execute() {
	c.receiver.Operation()
}

// Invoker класс, который вызывает команду.
type Invoker struct {
	command Command
}

// SetCommand назначает команду.
func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

// ExecuteCommand вызов команды.
func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func main() {
	receiver := &Receiver{}
	command := NewConcreteCommand(receiver)
	invoker := &Invoker{}
	invoker.SetCommand(command)
	invoker.ExecuteCommand()
}
