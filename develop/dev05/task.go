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
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-F - "fixed", точное совпадение со строкой, не паттерн
-i - "ignore-case" (игнорировать регистр)
-n - "line num", печатать номер строки
-c - "count" (количество строк) true
-v - "invert" (вместо совпадения, исключать) true
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flags struct {
	A  bool
	B  bool
	C  bool
	c  bool
	i  bool
	v  bool
	F  bool
	n  bool
	an int
	bn int
}

func main() {
	gp := &flags{}
	gp.grep()
}

func (f *flags) grep() {
	if len(os.Args) > 1 {
		f.parsFlags(os.Args)
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

	data := openfile(file)

	f.myfind(pattern, data)
}

func (f *flags) myfind(patt string, fl []string) {
	var rg *regexp.Regexp
	var err error
	if f.F {
		rg = func(s string) *regexp.Regexp {
			return regexp.MustCompile("^" + s + "$")
		}(patt)
	} else {
		rg, err = regexp.Compile(patt)
		if err != nil {
			log.Fatal(err)
		}
	}
	contextline := []int{}
	count, linecount := 0, 1
	for _, i := range fl {
		ok := rg.MatchString(i)
		if f.i {
			if strings.Contains(strings.ToLower(i), strings.ToLower(patt)) {
				fmt.Println(i)
			}
		} else if ok {
			if f.c {
				count++
			}
			if f.n {
				fmt.Printf("%v:%v\n", linecount, i)
			}
			if f.F {
				fmt.Println(i)
			}
		} else if !ok && f.v {
			fmt.Println(i)
		}
		if ok && (f.C || f.B || f.A) {
			contextline = append(contextline, linecount)
		}
		linecount++
	}
	if f.c {
		fmt.Println(count)
	}

	if f.C || f.B || f.A {
		f.contextABC(fl, contextline)
	}
}

func (f *flags) contextABC(fl []string, c []int) {
	j, k := 0, 0
	for indx, i := range c {
		if f.A {
			f.configA(&j, &k, &i, len(fl))
		} else if f.B {
			f.configB(&j, &k, &i)
		} else if f.C {
			f.configC(&j, &k, &i, len(fl))
		}
		for j < k {
			fmt.Println(fl[j])
			j++
		}
		if indx < len(c)-1 {
			fmt.Println()
		}
	}
}
func (f *flags) convert(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}
	return n
}

func (f *flags) configA(j *int, k *int, i *int, ln int) {
	*j = *i - 1
	if *i < 0 {
		*i = 0
	}
	*k = *i + f.an
	if *k > ln {
		*k = ln
	}
}

func (f *flags) configB(j *int, k *int, i *int) {
	*j = *i - f.bn - 1
	if *j < 0 {
		*j = 0
	}
	*k = *i
}

func (f *flags) configC(j *int, k *int, i *int, ln int) {
	*j = *i - f.bn - 1
	if *j < 0 {
		*j = 0
	}

	*k = *i + f.an
	if *k > ln {
		*k = ln
	}
}

func openfile(fl string) []string {
	file, err := os.Open(fl)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	sz, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	data := make([]string, 0, sz.Size())
	bf := bufio.NewScanner(file)
	for bf.Scan() {
		data = append(data, bf.Text())
	}

	return data
}
