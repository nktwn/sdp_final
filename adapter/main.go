package main

import "fmt"

// Target интерфейс, который ожидает клиент.
type Target interface {
	Request() string
}

// Adaptee класс, который нужно адаптировать.
type Adaptee struct{}

// SpecificRequest метод Adaptee, интерфейс которого несовместим с Target.
func (a *Adaptee) SpecificRequest() string {
	return "Specific request"
}

// Adapter адаптирует Adaptee к Target.
type Adapter struct {
	adaptee *Adaptee
}

// Request реализация метода интерфейса Target.
func (a *Adapter) Request() string {
	// Адаптер переводит вызовы из одного интерфейса в другой.
	return a.adaptee.SpecificRequest()
}

// Клиентский код, ожидающий работать с объектами, реализующими интерфейс Target.
func clientCode(target Target) {
	fmt.Println(target.Request())
}

func main() {
	// Клиент хочет использовать интерфейс Target.
	adaptee := &Adaptee{}
	adapter := &Adapter{adaptee: adaptee}

	// Мы можем использовать адаптер, как если бы это был Target.
	clientCode(adapter)
}
