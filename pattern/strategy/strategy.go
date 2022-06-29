package main

// Стратегия определяет общий интерфейс для всех вариаций алгоритма
// Контекст использует этот интерфейс для вызова алгоритма.
type RouteStrategy interface {
	buildRoute(from int, to int)
}
