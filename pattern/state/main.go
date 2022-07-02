/*
Состояние предлагает создать отдельные классы для каждого состояния,
в котором может пребывать объект, а затем вынести туда поведения,
соответствующие этим состояниям.

Применимость:
* поведение которого кардинально меняется в зависимости от внутреннего
	состояния, причём типов состояний много, и их код часто меняется.
* код класса содержит множество больших,
похожих друг на друга, условных операторов, которые выбирают поведения в зависимости от текущих значений полей класса.
* Когда сознательно используете табличную машину состояний,
построенную на условных операторах, но миритесь
с дублированием кода для похожих состояний и переходов

+ избавляет от множества больших условных операторов
+ концентрирует в одном месте код, связанный с определенным состоянием
+ упрощает код контекста

- может неоправданно усложнить код, если состояний мало и они редко меняются

+

*/

package main

import (
	"fmt"
	"log"
)

func main() {
	vendingMachine := NewVendingMachine(1, 10)

	err := vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.AddItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
