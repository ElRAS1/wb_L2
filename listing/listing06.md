Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
3 2 3

в го в функцию по умолчанию передаем копию, так как слайс это структура с указателем на базовый массив мы передадим копию указателя, а ответ будет таким потому что append возвращает новый слайс если мы выйдем за емкость(capacity) слайса. так как мы при создании слайса явно не указали емкость слайса он будет равен len который равен 3. Итого после append у нас уже будет новый слайс который будет указывать уже на другую ячейку памяти. а 1 поменяться на 3 из за того что до append успели поменять первое значение.
```
...

```
