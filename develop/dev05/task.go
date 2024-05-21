package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:

-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-F - "fixed", точное совпадение со строкой, не паттерн

-i - "ignore-case" (игнорировать регистр)
-A - "after" печатать +N строк после совпадения
-n - "line num", печатать номер строки
-c - "count" (количество строк) true
-v - "invert" (вместо совпадения, исключать) true
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flags struct {
	A     bool
	B     bool
	C     bool
	c     bool
	i     bool
	v     bool
	F     bool
	n     bool
	flagA bool
	an    int
	bn    int
}

func main() {
	gp := &flags{}
	gp.grep()
}

func (f *flags) grep() {
	fl := false
	if len(os.Args) > 1 {
		f.parsFlags(os.Args)
	} else {
		fl = true
		_ = fl
	}

}

func (f *flags) parsFlags(arg []string) {
	fl := arg[1]
	switch fl {
	case "A":
		f.A = true
		f.an = f.convert(os.Args[2])
	case "B":
		f.B = true
		f.bn = f.convert(os.Args[2])
	case "C":
		f.C = true
		f.an = f.convert(os.Args[2])
		f.bn = f.an
	case "c":
		f.c = true
	case "i":
		f.i = true
	case "v":
		f.v = true
	case "F":
		f.F = true
	case "n":
		f.n = true
	default:
		log.Fatalln("Undefined options...")
	}
	pattern := arg[2]
	file := arg[3]
	if f.C || f.A || f.B {
		pattern = arg[3]
		file = arg[4]
	}
	f.myfind(pattern, file)
}

func (f *flags) myfind(patt string, fl string) {
	file, err := os.Open(fl)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	text := bufio.NewScanner(file)
	rg, err := regexp.Compile(patt)

	if err != nil {
		log.Fatal(err)
	}
	count, linecount := 0, 1
	for text.Scan() {
		tmp := text.Text()
		ok := rg.MatchString(tmp)
		if f.i {
			if strings.Contains(strings.ToLower(tmp), strings.ToLower(patt)) {
				fmt.Println(tmp)
			}
		}
		if ok || f.flagA {
			f.Ok(tmp, patt, &count, &linecount)
		} else if !ok && f.v {
			fmt.Println(tmp)
		}
		linecount++
	}
	if f.c {
		fmt.Println(count)
	}
}

func (f *flags) Ok(tmp string, pat string, count *int, linecount *int) {
	if f.c {
		*count++
	}
	if f.n {
		fmt.Printf("%v:%v\n", *linecount, tmp)
	}
	if f.A && f.an >= 0 {
		f.flagA = true
		fmt.Println(tmp)
	}
	f.an--
	f.bn--
}

// func (f *flags) ContextandB() {

// }

func (f *flags) convert(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}
	return n
}
