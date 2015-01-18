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
	Execute()
	CheckParams() bool
}

func Factory(Drive filesystem.Drive, Outputter console.Outputter) []ICommand {
	var commands []ICommand

	commands = append(commands, CmdMkDir{Command{"mkdir", Drive, Outputter}})
	commands = append(commands, CmdMkFile{Command{"mkfile", Drive, Outputter}})

	return commands
}
