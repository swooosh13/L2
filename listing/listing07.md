Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Сначала будет выводить значения из каналов как fan-in
После будут идти дефолтные значения (v), т.к. канал закрыт, происходит livelock.

Варианты решения:
1. принимать 2 параметра (v, ok) в select case и проверять что канал закрыт\нет
2. Сделать в main горутине вместо range обычный for до N всех передаваемых заначений в каналы. Т.к. для остановки range нужно чтобы канал был закрыт

...

```
