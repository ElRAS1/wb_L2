package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern


Преимущества
	Позволяет изменять внутреннее представление продукта.
	Инкапсулирует код для построения и представления.
	Обеспечивает контроль над этапами процесса строительства.
	Недостатки
К недостаткам шаблона Builder относятся:
	Для каждого типа изделия должен быть создан отдельный конструктор ConcreteBuilder.
	Классы Builder должны быть изменяемыми.
	Может затруднять внедрение зависимостей.
*/

type Computer struct {
	Monitor int
	Mouse   string
	Cpu     string
}

type ComputerBuilder struct {
	monitor int
	mouse   string
	cpu     string
}

func NewComputer() *ComputerBuilder {
	return &ComputerBuilder{}
}

type ComputerI interface {
	SetMonitor(mon int) ComputerI
	SetMouse(mouse string) ComputerI
	SetCpu(cpu string) ComputerI

	Build() *Computer
}

func (c ComputerBuilder) SetMonitor(mon int) ComputerI {
	c.monitor = mon
	return c
}

func (c ComputerBuilder) SetMouse(mouse string) ComputerI {
	c.mouse = mouse
	return c
}

func (c ComputerBuilder) SetCpu(cpu string) ComputerI {
	c.cpu = cpu
	return c
}

func (c ComputerBuilder) Build() *Computer {
	return &Computer{Monitor: c.monitor, Mouse: c.mouse, Cpu: c.cpu}
}

func main() {
	pc := NewComputer()

	computer := pc.SetCpu("amd").SetMonitor(1).SetMouse("logitec").Build()

	fmt.Println(*computer)
}
