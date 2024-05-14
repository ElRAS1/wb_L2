package main

import (
	"testing"
)

func TestStr1(t *testing.T) {
	first := "a4bc2d5e"
	expected := "aaaabccddddde"
	res, err := Unpacking(first)
	if err != nil {
		t.Log("func returning error...")
	}
	if res != expected {
		t.Errorf("Expected %v, got %v", expected, res)
	}
	t.Logf("%v -> %v\n", first, res)
}

func TestStr2(t *testing.T) {
	first := "abcd"
	expected := "abcd"
	res, err := Unpacking(first)
	if err != nil {
		t.Log("func returning error...")
	}
	if res != expected {
		t.Errorf("Expected %v, got %v", expected, res)
	}
	t.Logf("%v -> %v\n", first, res)
}

func TestStr3(t *testing.T) {
	first := "45"
	expected := ""
	res, err := Unpacking(first)

	if err != nil {
		t.Logf("%v", err)
	}
	if res != expected {
		t.Errorf("Expected %v, got %v", expected, res)
	}
	t.Logf("bad string = %v", first)
}

func TestStr4(t *testing.T) {
	first := ""
	expected := ""
	res, err := Unpacking(first)
	if err != nil {
		t.Log("func returning error...")
	}
	if res != expected {
		t.Errorf("Expected %v, got %v", expected, res)
	}
	t.Logf("%v -> %v\n", first, res)
}
