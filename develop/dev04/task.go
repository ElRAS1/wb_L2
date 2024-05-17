package main

import (
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type ByRune []rune

// реализуем методы для sort.Slice
func (b ByRune) Len() int {
	return len(b)
}

func (b ByRune) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByRune) Less(i, j int) bool {
	return b[i] < b[j]
}

func Anagram(m *[]string) map[string][]string {

	toLower(*m)
	*m = remDuplicate(*m)
	// создаем карту для результата
	res := make(map[string][]string)

	for _, words := range *m {
		// приводим строку к рунам и сортируем его
		var tmp = ByRune(words)
		sort.Sort(tmp)
		// проверка на то что отсортированное слово уже есть в мапе
		_, ok := res[string(tmp)]

		if !ok {
			// если нет то ключом становиться отсортированное слово и его изначальное состояние добавляем в значение
			res[string(tmp)] = []string{words}

		} else {
			// по остортиванной строке добавялем изначальное состояние строки
			res[string(tmp)] = append(res[string(tmp)], words)

		}
	}

	return res
}

// приводим к нижнему регистру
func toLower(m []string) {
	for indx := range m {
		m[indx] = strings.ToLower(m[indx])
	}
}

func remDuplicate(m []string) []string {
	duplMap := make(map[string]struct{})

	for indx, i := range m {
		_, ok := duplMap[i]

		if !ok {
			duplMap[i] = struct{}{}
		} else {
			m = append(m[:indx], m[indx+1:]...)
		}

	}

	return m
}
