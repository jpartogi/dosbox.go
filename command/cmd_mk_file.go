package command

type CmdMkFile struct {
	Command
}

func (c CmdMkFile) CheckParams() bool {
	return true
}

func (c CmdMkFile) GetName() string {
	return c.Name
}

func (c CmdMkFile) Execute() {

}
