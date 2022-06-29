/*
Применение когда:
* Выполнить операцию над всеми элементами сложной структуры объектов
* когда поведение имеет смысл только для некоторых классов из существующей иерархии.

+ упрощает поведение операций, работающих со сложными структурами объектов
+ объединяет родственные операции в одном классе
+ посетитель может накапливать состояние при обходе структуры элементов

- если иерархия часто меняется - паттерн не оправдывает себя
- может привести к нарушению инкапсуляции эл-тов
*/

package main

import "fmt"

func main() {
	square := &square{
		side: 4,
	}

	circle := &circle{
		radius: 3,
	}

	rec := &rectangle{
		l: 2,
		b: 3,
	}

	areaCalculator := &areaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rec.accept(areaCalculator)

	fmt.Println(areaCalculator.area)

	middleCoordinates := &middleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)

	fmt.Println(middleCoordinates)
}
