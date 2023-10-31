package main

import "fmt"

type Command interface {
	Execute()
}

type Receiver struct {
	Name string
}

func (r *Receiver) Action() {
	fmt.Printf("%s is performing an action\n", r.Name)
}

type ConcreteCommand struct {
	receiver *Receiver
}

func (cc *ConcreteCommand) Execute() {
	cc.receiver.Action()
}

type Invoker struct {
	commands []Command
}

func (i *Invoker) AddCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *Invoker) ExecuteCommands() {
	for _, command := range i.commands {
		command.Execute()
	}
}

func main() {
	receiver := &Receiver{Name: "Receiver 1"}

	command1 := &ConcreteCommand{receiver}
	command2 := &ConcreteCommand{receiver}
	command3 := &ConcreteCommand{receiver}

	invoker := &Invoker{}
	invoker.AddCommand(command1)
	invoker.AddCommand(command2)
	invoker.AddCommand(command3)

	invoker.ExecuteCommands()
}
