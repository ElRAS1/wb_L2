Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
вывод будет от 0 до 9

в анонимной функции которая будет выполняться в отдельной горутин булет запущен цикл от 0 до 9 и каждое значение i будет записываться в канал ch после каждой итерации канал будет блокироваться для записи так как он не буферизированный и будет ждать пока из него прочитают. в это время в main в цикле будет читаться из канала один и так будет при каждой итерации сначала запишем в канал(заблокируемся) и прочитаем из канала(после прочтение канал заблокиреться для записи и разблокиреться для записи)

в конце когда цикл дойдет до 9 мы запишем в канал и заблокируемся, в main прочитаем из канала и тоже заблокируемся. а так как цикл завершиться и потом не будет закрытия канала мы словим deadlock.
```
...

```
