package main

import "fmt"

type Subscriber interface {
	Update(context *Publisher)
}

type Publisher struct {
	subscribers []Subscriber
	mainState   int
}

func (p *Publisher) Subscribe(s Subscriber) {
	p.subscribers = append(p.subscribers, s)
}

func (p *Publisher) NotifySubscribers() {
	for _, s := range p.subscribers {
		s.Update(p)
	}
}

func (p *Publisher) SetState(state int) {
	p.mainState = state
	p.NotifySubscribers()
}

type ConcreteSubscriber struct {
	id int
}

func (s *ConcreteSubscriber) Update(context *Publisher) {
	fmt.Printf("Subscriber %d received update. New state: %d\n", s.id, context.mainState)
}

func main() {
	publisher := &Publisher{}
	subscriber1 := &ConcreteSubscriber{id: 1}
	publisher.Subscribe(subscriber1)
	publisher.SetState(10)
}
