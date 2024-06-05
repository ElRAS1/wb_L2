package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern

	Фасад — это структурный паттерн проектирования, который предоставляет простой интерфейс к сложной системе классов, библиотеке или фреймворку.
	Фасад — это простой интерфейс для работы со сложной подсистемой, содержащей множество классов. Фасад может иметь урезанный интерфейс, не имеющий 100% функциональности,
	которой можно достичь, используя сложную подсистему напрямую. Но он предоставляет именно те фичи, которые нужны клиенту, и скрывает все остальные.
	Фасад предоставляет унифицированный интерфейс вместо набора интерфейсов некоторой подсистемы. Фасад определяет интерфейс более высокого уровня, который упрощает использование подсистемы.

	 + Изолирует клиентов от компонентов сложной подсистемы.
	 - Фасад рискует стать божественным объектом, привязанным ко всем классам программы.
*/

// Структура для компьютера
type Facade_computer struct {
	ram  *Ram
	cpu  *CPU
	disc *HardDisk
}

// Создание объекта компьютера
func NewFacade_computer() *Facade_computer {
	return &Facade_computer{&Ram{}, &CPU{}, &HardDisk{}}
}

// Структура для CPU
type CPU struct{}

// Метод для работы с CPU
func (c CPU) CpuProcess() {
	fmt.Println("CPU operation")
}

// Структура для RAM
type Ram struct{}

// Метод для работы с RAM
func (r Ram) RamProcess() {
	fmt.Println("Ram operation")
}

// Структура для HardDisk
type HardDisk struct {
}

// Метод для работы с HardDisk
func (h HardDisk) DiskProcess() {
	fmt.Println("Disk operation")
}

func main() {

	computer := NewFacade_computer()

	computer.cpu.CpuProcess()
	computer.disc.DiskProcess()
	computer.ram.RamProcess()

}
