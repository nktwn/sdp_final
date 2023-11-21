package main

import "fmt"

// Observer интерфейс, который должен реализовать каждый наблюдатель.
type Observer interface {
	Update(subject *Subject)
}

// Subject структура, за которой будут следить наблюдатели.
type Subject struct {
	state     int
	observers []Observer
}

// SetState устанавливает состояние субъекта и уведомляет наблюдателей.
func (s *Subject) SetState(state int) {
	s.state = state
	s.notifyObservers()
}

// Attach добавляет нового наблюдателя к субъекту.
func (s *Subject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

// notifyObservers уведомляет всех наблюдателей о изменении состояния.
func (s *Subject) notifyObservers() {
	for _, observer := range s.observers {
		observer.Update(s)
	}
}

// ConcreteObserver конкретный наблюдатель, который реагирует на уведомления.
type ConcreteObserver struct {
	id int
}

// Update реализация метода Update для ConcreteObserver.
func (o *ConcreteObserver) Update(subject *Subject) {
	fmt.Printf("Observer %d: subject has new state %d\n", o.id, subject.state)
}

func main() {
	// Создание субъекта
	subject := &Subject{}

	// Создание и добавление наблюдателей
	observer1 := &ConcreteObserver{id: 1}
	observer2 := &ConcreteObserver{id: 2}

	subject.Attach(observer1)
	subject.Attach(observer2)

	// Изменение состояния субъекта, что приведет к уведомлению наблюдателей
	subject.SetState(10)
	subject.SetState(20)
}
