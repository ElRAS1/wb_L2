package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Unpacking(s string) (string, error) {
	if _, err := strconv.Atoi(s); err == nil {
		return "", fmt.Errorf("(некорректная строка)")
	}

	if s == "" {
		return "", nil
	}
	var res strings.Builder
	ch := ""
	for _, i := range s {

		if ok := unicode.IsDigit(i); !ok {
			res.WriteString(string(i))
			ch = string(i)
		} else {
			r, err := strconv.Atoi(string(i))
			if err != nil {
				log.Fatalln(err)
			}

			for i := 0; i < r-1; i++ {
				res.WriteString(ch)
			}
		}
	}
	return res.String(), nil
}
