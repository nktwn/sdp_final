package main

import "fmt"

// Subject интерфейс, определяющий общие операции для RealSubject и Proxy.
type Subject interface {
	Request()
}

// RealSubject реальный объект, операции которого нужно контролировать.
type RealSubject struct{}

func (s *RealSubject) Request() {
	fmt.Println("RealSubject: Обработка запроса.")
}

// Proxy контролирует доступ к объекту RealSubject.
type Proxy struct {
	realSubject *RealSubject
}

func (p *Proxy) Request() {
	if p.realSubject == nil {
		fmt.Println("Proxy: Создание реального объекта.")
		p.realSubject = &RealSubject{}
	}

	// Передача вызова реальному объекту
	fmt.Println("Proxy: Логика до вызова.")
	p.realSubject.Request()
	fmt.Println("Proxy: Логика после вызова.")
}

func main() {
	proxy := &Proxy{}

	// Клиент работает с объектом через прокси.
	proxy.Request()
}
