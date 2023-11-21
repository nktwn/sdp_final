package main

import "fmt"

// Button интерфейс для продукта типа "Кнопка".
type Button interface {
	Paint() string
}

// WinButton конкретная реализация кнопки для Windows.
type WinButton struct{}

func (b *WinButton) Paint() string {
	return "Windows Button"
}

// MacButton конкретная реализация кнопки для MacOS.
type MacButton struct{}

func (b *MacButton) Paint() string {
	return "MacOS Button"
}

// GUIFactory интерфейс абстрактной фабрики для создания элементов GUI.
type GUIFactory interface {
	CreateButton() Button
}

// WinFactory конкретная фабрика для создания элементов GUI для Windows.
type WinFactory struct{}

func (f *WinFactory) CreateButton() Button {
	return &WinButton{}
}

// MacFactory конкретная фабрика для создания элементов GUI для MacOS.
type MacFactory struct{}

func (f *MacFactory) CreateButton() Button {
	return &MacButton{}
}

// Application структура, использующая абстрактную фабрику.
type Application struct {
	factory GUIFactory
}

func (a *Application) CreateUI() {
	button := a.factory.CreateButton()
	fmt.Println(button.Paint())
}

func main() {
	var factory GUIFactory

	// Конфигурация для Windows
	factory = &WinFactory{}
	app := Application{factory: factory}
	app.CreateUI()

	// Конфигурация для MacOS
	factory = &MacFactory{}
	app = Application{factory: factory}
	app.CreateUI()
}
