package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	first := []string{"1 2 3 4 5"}
	expected := "5 4 3 2 1"
	res := ReversSort(first)

	if res[0] != expected {
		t.Errorf("Expected %v, got %v", expected, res[0])
	}
	t.Logf("%v -> %v\n", first, res)
}

func TestUniq(t *testing.T) {
	first := []string{"a", "b", "a"}
	expected := []string{"a", "b"}
	res := Uniq(first)

	for i := 0; i < len(expected)-1; i++ {
		if res[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, res[0])
			return
		}
	}
	t.Logf("%v -> %v\n", first, res)
}

func TestK(t *testing.T) {
	first := []string{"51 1", "9 6", "8 0"}
	expected := []string{"8 0", "51 1", "9 6"}
	res := SortByColumn(first, 2)

	for i := 0; i < len(expected)-1; i++ {
		if res[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, res[0])
			return
		}
	}
	t.Logf("%v -> %v\n", first, res)
}

func TestN(t *testing.T) {
	first := []string{"51", "45", "99"}
	expected := []string{"45", "51", "99"}
	res := Nsort(first)

	for i := 0; i < len(expected)-1; i++ {
		if res[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, res[0])
			return
		}
	}
	t.Logf("%v -> %v\n", first, res)
}
