package main

// Контекст всегда работает со стратегиями через обзий интерфейс.
// Не знает какая именно стратегия ему подана
type Navigator struct {
	strategy RouteStrategy
}

func (n *Navigator) SetStrategy(strategy RouteStrategy) {
	n.strategy = strategy
}

func (n *Navigator) Route(from int, to int) {
	n.strategy.buildRoute(from, to)
}


