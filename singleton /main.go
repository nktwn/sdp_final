package main

import (
	"fmt"
	"sync"
)

// Singleton представляет собой структуру, которая будет реализовывать паттерн Одиночка.
type Singleton struct {
	// Данные, которые будут внутри одиночки.
	Value int
}

var (
	// instance - экземпляр Singleton, который будет использоваться.
	instance *Singleton
	// once - используется для выполнения инициализации один раз.
	once sync.Once
)

// GetInstance возвращает экземпляр Singleton, инициализируя его при первом вызове.
func GetInstance() *Singleton {
	// Инициализация происходит только один раз
	once.Do(func() {
		instance = &Singleton{Value: 42}
	})
	return instance
}

func main() {
	// Первый вызов GetInstance, который инициализирует экземпляр
	singleton := GetInstance()
	fmt.Println(singleton.Value) // Выведет 42

	// Повторные вызовы GetInstance возвращают тот же экземпляр
	sameSingleton := GetInstance()
	fmt.Println(sameSingleton.Value) // Также выведет 42

	// Изменение значения в одном экземпляре
	singleton.Value = 1

	// Изменения отразятся и в другом экземпляре, так как они ссылаются на один и тот же объект
	fmt.Println(sameSingleton.Value) // Теперь выведет 1
}
