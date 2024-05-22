package main

import (
	"bytes"
	"os/exec"
	"testing"
)

// -A - "after" печатать +N строк после совпадения
// -B - "before" печатать +N строк до совпадения
// -C - "context" (A+B) печатать ±N строк вокруг совпадения
// -F - "fixed", точное совпадение со строкой, не паттерн
// -i - "ignore-case" (игнорировать регистр)
// -n - "line num", печатать номер строки
// -c - "count" (количество строк) true
// -v - "invert" (вместо совпадения, исключать)
func TestOptionsC(t *testing.T) {
	expectedCmd := exec.Command("grep", "-c", "hello", "test.txt")
	var buf bytes.Buffer
	expectedCmd.Stdout = &buf
	err := expectedCmd.Run()
	if err != nil {
		t.Fatalf("Ошибка выполнения команды grep: %v", err)
	}

	actualCmd := exec.Command("go", "run", "task.go", "c", "hello", "test.txt")
	var actualBuf bytes.Buffer
	actualCmd.Stdout = &actualBuf
	err = actualCmd.Run()
	if err != nil {
		t.Fatalf("Ошибка выполнения моей программы: %v", err)
	}

	// Сравниваем выводы двух команд
	if buf.String() != actualBuf.String() {
		t.Errorf("Ожидаемый вывод:\n%v\nИтоговый вывод:\n%v", buf.String(), actualBuf.String())
	}
}

func TestOptionsV(t *testing.T) {
	expectedCmd := exec.Command("grep", "-v", "hello", "test.txt")
	var buf bytes.Buffer
	expectedCmd.Stdout = &buf
	err := expectedCmd.Run()
	if err != nil {
		t.Fatalf("Ошибка выполнения команды grep: %v", err)
	}

	actualCmd := exec.Command("go", "run", "task.go", "v", "hello", "test.txt")
	var buf2 bytes.Buffer
	actualCmd.Stdout = &buf2
	err2 := actualCmd.Run()
	if err2 != nil {
		t.Fatalf("Ошибка выполнения моей команды grep: %v", err2)
	}

	if buf.String() != buf2.String() {
		t.Errorf("Ожидаемый вывод:\n%v\nИтоговый вывод:\n%v", buf.String(), buf2.String())
	}
}

func TestOptionsN(t *testing.T) {
	expectedCmd := exec.Command("grep", "-n", "hello", "test.txt")
	var buf bytes.Buffer
	expectedCmd.Stdout = &buf
	err := expectedCmd.Run()
	if err != nil {
		t.Fatalf("Ошибка выполнения команды grep: %v", err)
	}

	actualCmd := exec.Command("go", "run", "task.go", "n", "hello", "test.txt")
	var buf2 bytes.Buffer
	actualCmd.Stdout = &buf2
	err2 := actualCmd.Run()
	if err2 != nil {
		t.Fatalf("Ошибка выполнения моей команды grep: %v", err2)
	}

	if buf.String() != buf2.String() {
		t.Errorf("Ожидаемый вывод:\n%v\nИтоговый вывод:\n%v", buf.String(), buf2.String())
	}
}

func TestOptionsI(t *testing.T) {
	expectedCmd := exec.Command("grep", "-i", "hello", "test.txt")
	var buf bytes.Buffer
	expectedCmd.Stdout = &buf
	err := expectedCmd.Run()
	if err != nil {
		t.Fatalf("Ошибка выполнения команды grep: %v", err)
	}

	actualCmd := exec.Command("go", "run", "task.go", "i", "hello", "test.txt")
	var buf2 bytes.Buffer
	actualCmd.Stdout = &buf2
	err2 := actualCmd.Run()
	if err2 != nil {
		t.Fatalf("Ошибка выполнения моей команды grep: %v", err2)
	}

	if buf.String() != buf2.String() {
		t.Errorf("Ожидаемый вывод:\n%v\nИтоговый вывод:\n%v", buf.String(), buf2.String())
	}
}

func TestOptionsContext(t *testing.T) {
	expectedCmd := exec.Command("grep", "-C", "3", "cat", "test.txt")
	var buf bytes.Buffer
	expectedCmd.Stdout = &buf
	err := expectedCmd.Run()
	if err != nil {
		t.Fatalf("Ошибка выполнения команды grep: %v", err)
	}

	actualCmd := exec.Command("go", "run", "task.go", "C", "3", "cat", "test.txt")
	var buf2 bytes.Buffer
	actualCmd.Stdout = &buf2
	err2 := actualCmd.Run()
	if err2 != nil {
		t.Fatalf("Ошибка выполнения моей команды grep: %v", err2)
	}

	if buf.String() != buf2.String() {
		t.Errorf("Ожидаемый вывод:\n%v\nИтоговый вывод:\n%v", buf.String(), buf2.String())
	}
}

func TestOptionsA(t *testing.T) {
	expectedCmd := exec.Command("grep", "-A", "3", "cat", "test.txt")
	var buf bytes.Buffer
	expectedCmd.Stdout = &buf
	err := expectedCmd.Run()
	if err != nil {
		t.Fatalf("Ошибка выполнения команды grep: %v", err)
	}

	actualCmd := exec.Command("go", "run", "task.go", "A", "3", "cat", "test.txt")
	var buf2 bytes.Buffer
	actualCmd.Stdout = &buf2
	err2 := actualCmd.Run()
	if err2 != nil {
		t.Fatalf("Ошибка выполнения моей команды grep: %v", err2)
	}

	if buf.String() != buf2.String() {
		t.Errorf("Ожидаемый вывод:\n%v\nИтоговый вывод:\n%v", buf.String(), buf2.String())
	}
}

func TestOptionsB(t *testing.T) {
	expectedCmd := exec.Command("grep", "-B", "3", "cat", "test.txt")
	var buf bytes.Buffer
	expectedCmd.Stdout = &buf
	err := expectedCmd.Run()
	if err != nil {
		t.Fatalf("Ошибка выполнения команды grep: %v", err)
	}

	actualCmd := exec.Command("go", "run", "task.go", "B", "3", "cat", "test.txt")
	var buf2 bytes.Buffer
	actualCmd.Stdout = &buf2
	err2 := actualCmd.Run()
	if err2 != nil {
		t.Fatalf("Ошибка выполнения моей команды grep: %v", err2)
	}

	if buf.String() != buf2.String() {
		t.Errorf("Ожидаемый вывод:\n%v\nИтоговый вывод:\n%v", buf.String(), buf2.String())
	}
}

func TestOptionsF(t *testing.T) {
	expectedCmd := exec.Command("grep", "-F", "cat", "test.txt")
	var buf bytes.Buffer
	expectedCmd.Stdout = &buf
	err := expectedCmd.Run()
	if err != nil {
		t.Fatalf("Ошибка выполнения команды grep: %v", err)
	}

	actualCmd := exec.Command("go", "run", "task.go", "F", "cat", "test.txt")
	var buf2 bytes.Buffer
	actualCmd.Stdout = &buf2
	err2 := actualCmd.Run()
	if err2 != nil {
		t.Fatalf("Ошибка выполнения моей команды grep: %v", err2)
	}

	if buf.String() != buf2.String() {
		t.Errorf("Ожидаемый вывод:\n%v\nИтоговый вывод:\n%v", buf.String(), buf2.String())
	}
}
