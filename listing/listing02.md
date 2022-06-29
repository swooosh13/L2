Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
defer выполняются с последнего к первому (как LIFO)

1. При именнованом возврате  
	1. 1. Именованная возвращаемая перменная назначается 
	1. 2. defer увеличивает ее 
x = 1 
x++
x == 2 - конечный результат именованной переменной

2. Во втором случае все как обычно
Вывод - 1
...

```
