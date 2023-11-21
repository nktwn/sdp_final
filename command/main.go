package main

import "fmt"

// Command интерфейс команды, который объявляет метод для выполнения команды.
type Command interface {
	Execute()
}

// ConcreteCommand конкретная реализация команды, которая выполняет действие.
type ConcreteCommand struct {
	receiver *Receiver
}

// Execute выполняет команду, делегируя вызов получателю.
func (c *ConcreteCommand) Execute() {
	c.receiver.Action()
}

// Receiver класс, который знает, как выполнять операции.
type Receiver struct {
}

// Action операция, выполняемая получателем.
func (r *Receiver) Action() {
	fmt.Println("Receiver: выполнение действия")
}

// Invoker держатель и инициатор команды.
type Invoker struct {
	command Command
}

// SetCommand назначает команду.
func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

// ExecuteCommand выполняет команду.
func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func main() {
	receiver := &Receiver{}
	command := &ConcreteCommand{receiver}
	invoker := &Invoker{}

	invoker.SetCommand(command)
	invoker.ExecuteCommand()
}
