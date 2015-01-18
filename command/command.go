package command

import (
	"github.com/jpartogi/dosboxgo/console"
	"github.com/jpartogi/dosboxgo/filesystem"
)

type Command struct {
	Name      string
	Drive     filesystem.Drive
	Outputter console.Outputter
}

type ICommand interface {
	GetName() string
	Execute(Params []string)
	CheckParams(Params []string) error
}

func Factory(Drive filesystem.Drive, Outputter console.Outputter) []ICommand {
	var commands []ICommand

	// Add your commands here
	commands = append(commands, CmdCd{Command{"cd", Drive, Outputter}})
	commands = append(commands, CmdDir{Command{"dir", Drive, Outputter}})
	commands = append(commands, CmdMkDir{Command{"mkdir", Drive, Outputter}})
	commands = append(commands, CmdMkFile{Command{"mkfile", Drive, Outputter}})

	return commands
}
