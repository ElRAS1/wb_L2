package main

import (
	// "fmt"
	"reflect"
	"testing"
)

func TestFindAnagram(t *testing.T) {
	first := []string{"Пятак", "Листок", "Пятка", "Слиток", "Тяпка", "Столик", "Столик"}

	expectedKeys := []string{"акптя", "иклост"}                                               // Ключи, которые ожидаются в карте
	expectedValues := [][]string{{"пятак", "пятка", "тяпка"}, {"листок", "слиток", "столик"}} // Ожидаемые значения для каждого ключа

	res := Anagram(&first)

	// Проверяем, что количество элементов в результирующей карте соответствует ожидаемому
	if len(res) != len(expectedKeys) {
		t.Errorf("Unexpected number of elements in result map")
	}

	// Проверяем, что каждый ключ и соответствующие ему значения в карте соответствуют ожидаемым
	for i, key := range expectedKeys {
		if _, exists := res[key]; !exists {
			t.Errorf("Key %q not found in result map", key)
		} else if !reflect.DeepEqual(res[key], expectedValues[i]) {
			t.Errorf("Value for key %q does not match expected value", key)
		}
	}
}
