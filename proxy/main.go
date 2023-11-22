package main

import "fmt"

// ServiceInterface объявляет интерфейс сервиса.
type ServiceInterface interface {
	Operation()
}

// Service - реальный сервис, который мы хотим использовать.
type Service struct{}

// Operation - это метод Service, который мы хотим контролировать через прокси.
func (s *Service) Operation() {
	fmt.Println("Service operation executed.")
}

// Proxy - это структура прокси, которая имеет ссылку на сервис.
type Proxy struct {
	realService ServiceInterface
}

// NewProxy - конструктор для создания нового прокси.
func NewProxy(service ServiceInterface) *Proxy {
	return &Proxy{realService: service}
}

// checkAccess - это метод, который может выполнять проверки перед выполнением операции сервиса.
func (p *Proxy) checkAccess() bool {
	// Здесь могут быть некоторые проверки прав доступа.
	fmt.Println("Proxy check access.")
	return true
}

// Operation - метод прокси, который делегирует вызов реальному сервису.
func (p *Proxy) Operation() {
	if p.checkAccess() {
		p.realService.Operation()
	}
}

func main() {
	realService := &Service{}
	proxy := NewProxy(realService)
	proxy.Operation()
}
