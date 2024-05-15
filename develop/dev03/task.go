package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	// "sort"
	// "sort"
	// "strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	MySort()
}

func MySort() {

	k := flag.Int("k", 0, "enter column")
	n := flag.Bool("n", false, "sort by num")
	r := flag.Bool("r", false, "revers sort")
	u := flag.Bool("u", false, "do not dublicate lines")

	flag.Parse()

	str := ReadFile("test.txt")
	if *r {
		ReversSort(*str)
	} else if *u {
		Uniq(*str)
	} else if *n {
		Nsort(*str)
	} else if *k != -1 {
		SortByColumn(*str, *k)
	}

}

func ReversSort(str []string) {
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Print(str[i])
	}
}

func Uniq(str []string) {
	mp := make(map[string]struct{})

	for indx, _ := range str {
		_, ok := mp[str[indx]]
		if !ok {
			fmt.Print(str[indx])
		}
		mp[str[indx]] = struct{}{}
	}
}

func Nsort(str []string) {
	sort.Strings(str)
	for _, i := range str {
		fmt.Print(i)
	}
}

func ReadFile(s string) *[]string {
	f, _ := os.Open(s)
	defer f.Close()
	fileStats, err := f.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	str := make([]string, fileStats.Size())
	buf := bufio.NewScanner(f)

	for i := 0; buf.Scan(); i++ {
		str = append(str, buf.Text()+"\n")
	}

	return &str
}

func SortByColumn(data []string, colNum int) {
	sort.Slice(data, func(i, j int) bool {
		fields1 := strings.Fields(data[i])
		fields2 := strings.Fields(data[j])

		if len(fields1) > colNum && len(fields2) > colNum {
			val1, err1 := strconv.Atoi(fields1[colNum])
			val2, err2 := strconv.Atoi(fields2[colNum])

			if err1 == nil && err2 == nil {
				return val1 < val2
			}
		}

		return fields1[colNum] < fields2[colNum]
	})

	fmt.Print(data)
}
