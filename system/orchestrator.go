package system

import (
	"github.com/jpartogi/dosboxgo/console"
	"github.com/jpartogi/dosboxgo/filesystem"
)

func OrchestrateSystem() {
	drive := filesystem.NewDrive("C")
	drive.Restore()

	outputter := console.NewOutputter()

	proc := console.NewProc(drive, outputter)
	proc.ProcessInput()
}
