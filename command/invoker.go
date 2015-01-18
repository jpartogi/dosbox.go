package command

import (
	"fmt"
	"strings"
)

type Invoker struct {
	Commands []ICommand
}

func NewInvoker(Commands []ICommand) Invoker {
	return Invoker{Commands}
}

func (i *Invoker) SetCommands(Commands []ICommand) {
	i.Commands = Commands
}

func (i *Invoker) ExecuteCommand(Command string) {
	commandName, commandParams := ParseCommand(Command)

	for _, command := range i.Commands {
		if command.GetName() == commandName {

			error := command.CheckParams(commandParams)
			if error == nil {
				command.Execute(commandParams)
			} else {
				fmt.Println(error)
			}
		}
	}
}

func ParseCommand(Command string) (string, []string) {
	Command = strings.ToLower(strings.TrimSpace(Command))
	Command = strings.Replace(Command, ",", " ", -1)
	Command = strings.Replace(Command, ";", " ", -1)

	Commands := strings.Split(Command, " ")

	return Commands[0], Commands[1:]
}
