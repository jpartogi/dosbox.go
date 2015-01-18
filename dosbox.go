package main

import (
	"github.com/jpartogi/dosboxgo/command"
	"github.com/jpartogi/dosboxgo/console"
	"github.com/jpartogi/dosboxgo/filesystem"
)

func main() {
	drive := filesystem.NewDrive("C")
	drive.Restore()

	outputter := console.NewOutputter()

	commands := command.Factory(drive, outputter)

	invoker := command.NewInvoker(commands)

	proc := command.NewProc(invoker, drive, outputter)
	proc.ProcessInput()
}
