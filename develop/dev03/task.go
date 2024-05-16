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
	// "strings"
	// "sort"
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
		res := ReversSort(*str)
		PrintRes(res)
	} else if *u {
		res := Uniq(*str)
		PrintRes(res)
	} else if *n {
		res := Nsort(*str)
		PrintRes(res)
	} else if *k != -1 {
		res := SortByColumn(*str, *k)
		PrintRes(res)
	}

}

func ReversSort(str []string) []string {

	res := make([]string, len(str))
	for i, s := range str {
		rs := []rune(s)
		for j, k := 0, len(rs)-1; j < k; j, k = j+1, k-1 {
			rs[j], rs[k] = rs[k], rs[j]
		}
		res[i] = string(rs)
	}
	return res
}

func Uniq(str []string) []string {
	mp := make(map[string]struct{})
	res := make([]string, 0, len(str))
	for indx := range str {
		_, ok := mp[str[indx]]
		if !ok {
			res = append(res, str[indx])
			mp[str[indx]] = struct{}{}
		}
	}
	return res
}

func Nsort(str []string) []string {
	sort.Slice(str, func(i, j int) bool {
		// Преобразование строк в числа для сортировки
		iF, err1 := strconv.Atoi(str[i])
		jF, err2 := strconv.Atoi(str[j])
		if err1 != nil && err2 != nil {
			log.Fatalln(err1)
		}
		return iF < jF
	})

	return str
}

func ReadFile(s string) *[]string {
	f, _ := os.Open(s)
	defer f.Close()
	fileStats, err := f.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	str := make([]string, 0, fileStats.Size())
	buf := bufio.NewScanner(f)

	for i := 0; buf.Scan(); i++ {
		str = append(str, buf.Text())
	}

	sort.Strings(str)

	return &str
}

func SortByColumn(data []string, colNum int) []string {
	sort.Slice(data, func(i, j int) bool {
		iFields := strings.Split(data[i], " ") // Разделение строки на части по пробелам
		jFields := strings.Split(data[j], " ")

		if len(iFields)-1 > colNum && len(jFields)-1 > colNum {
			iF, err1 := strconv.Atoi(iFields[colNum])
			jF, err2 := strconv.Atoi(jFields[colNum])

			if err1 == nil && err2 == nil {
				return iF < jF
			}
		}
		return iFields[colNum] < jFields[colNum]
	})
	return data
}

func PrintRes(str []string) {
	for _, i := range str {
		fmt.Println(i)
	}
}
