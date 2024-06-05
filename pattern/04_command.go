package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern

Преимущества командного шаблона:
	Гибкость: Легко добавлять новые команды без изменения существующего кода.
	Логирование и отмена: Легко добавить логирование или возможность отмены операций.
	Транзакционность: Упрощает обработку транзакций, так как можно группировать команды вместе.
Недостатки командного шаблона:
	Затраты памяти: Создание объектов команды может привести к увеличению потребления памяти.
	Сложность: Шаблон может усложнить архитектуру приложения, если его использовать там, где это не требуется.
*/

type Command interface {
	Execute()
}

func NewActions(command Actions) *Actions {
	return &Actions{
		cmd: command.cmd,
	}
}

type Actions struct {
	cmd Command
}

func (a Actions) Start() {
	a.cmd.Execute()
}

///////////////////////////////////

type RunCommand struct {
	act Commands
}

func (f *RunCommand) Execute() {
	f.act.Running()
}

func NewRun(r Commands) *RunCommand {
	return &RunCommand{act: r}
}

type FlyCommand struct {
	act Commands
}

func (f *FlyCommand) Execute() {
	f.act.Fly()
}

func NewFly(r Commands) *FlyCommand {
	return &FlyCommand{act: r}
}

// ////////////////////////////////
type Commands interface {
	Running()
	Fly()
}

// ///////////////////////////////
type Man struct {
	action bool
}

func NewMan() *Man {
	return &Man{}
}

func (m Man) Running() {
	m.action = true
}

func (m Man) Fly() {
	m.action = true
}

// func main() {
// 	man := NewMan()
// 	man2 := NewMan()

// 	manRun := NewRun(man)
// 	manFly := NewFly(man2)

// 	cmd := NewActions(manRun)
// 	cmd2 := NewActions(manFly)
// }
