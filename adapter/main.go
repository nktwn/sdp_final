package main

import "fmt"

// ClientInterface описывает интерфейс, который клиент ожидает использовать.
type ClientInterface interface {
	Method(data string)
}

// Service класс, который имеет несовместимый интерфейс.
type Service struct{}

// ServiceMethod метод, который предоставляет сервис. Он имеет несовместимый интерфейс с клиентом.
func (s *Service) ServiceMethod(specialData string) {
	fmt.Println("ServiceMethod вызван с данными:", specialData)
}

// Adapter класс, который адаптирует сервис к клиентскому интерфейсу.
type Adapter struct {
	adaptedService *Service
}

// Method реализация клиентского интерфейса, который адаптирует вызов к сервисному методу.
func (a *Adapter) Method(data string) {
	// Тут может быть логика преобразования данных в формат, который требуется сервису.
	specialData := data // Простое присвоение для примера
	a.adaptedService.ServiceMethod(specialData)
}

// Клиентский код, который работает с объектами через клиентский интерфейс.
func main() {
	service := &Service{}
	adapter := &Adapter{adaptedService: service}
	clientUsesInterface(adapter, "тестовые данные")
}

// clientUsesInterface функция, показывающая, как клиент может использовать интерфейс.
func clientUsesInterface(clientInterface ClientInterface, data string) {
	clientInterface.Method(data)
}
