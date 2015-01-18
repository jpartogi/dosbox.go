package command

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
	for _, command := range i.Commands {
		if command.GetName() == Command {
			command.Execute()
		}
	}
}
