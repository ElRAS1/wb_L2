package main

import (
	"bufio"
	"flag"
	"fmt"

	// "fmt"
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
		CreateRes(res)
		PrintRes(res)
	} else if *u {
		res := Uniq(*str)
		CreateRes(res)
		PrintRes(res)
	} else if *n {
		res := Nsort(*str)
		CreateRes(res)
		PrintRes(res)
	} else if *k != -1 {
		res := SortByColumn(*str, *k)
		CreateRes(res)
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
	colNum--
	sort.Slice(data, func(i, j int) bool {
		iFields := strings.Fields(data[i])
		jFields := strings.Fields(data[j])
		if len(iFields) > colNum && len(jFields) > colNum {
			iF := iFields[colNum]
			jF := jFields[colNum]

			return iF < jF
		}
		return iFields[colNum] < jFields[colNum]
	})
	return data
}

func CreateRes(data []string) {
	file, err := os.Create("result.txt")
	if err != nil {
		panic("Cant create file")
	}
	defer file.Close()

	for i := 0; i < len(data); i++ {
		_, err := file.WriteString(data[i] + "\n")
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func PrintRes(str []string) {
	for _, i := range str {
		fmt.Println(i)
	}
}
