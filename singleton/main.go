package main

import (
	"fmt"
	"sync"
)

type Singleton struct{}

var (
	instance *Singleton
	once     sync.Once
)

func getInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	// Первый вызов getInstance, который инициализирует экземпляр
	singleton := getInstance()
	fmt.Printf("Instance: %p\n", singleton)

	// Повторные вызовы getInstance возвращают тот же экземпляр
	sameSingleton := getInstance()
	fmt.Printf("Instance: %p\n", sameSingleton)

	// Проверка, что обе переменные указывают на один и тот же экземпляр
	if singleton == sameSingleton {
		fmt.Println("Обе переменные содержат ссылку на один и тот же экземпляр.")
	}
}
