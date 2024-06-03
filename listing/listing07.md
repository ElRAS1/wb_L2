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
будут выводиться числа от 1 до 8, сначала запуститться функция asChan которая будет записывать в переменную a ф функции создаеться небуферизированный канал, это значит мы сможем записать одно значение и заблокируемся, потом в main опять запуститься asChan которая будет записывать в переменную b и так же блокироваться.


далее запуститься функция merge которая будет читать из и любого из двух каналов который будет доступен для чтения и печатать.

после того как в каналы перстанут записываться числа, мы будем бесконечно печатать 0 потому что не закрываем канал в функции merge.
```
...

```
