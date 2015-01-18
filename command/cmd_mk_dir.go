package command

type CmdMkDir struct {
	Command
}

func (c CmdMkDir) CheckParams() bool {
	return true
}

func (c CmdMkDir) GetName() string {
	return c.Name
}

func (c CmdMkDir) Execute() {
	c.Outputter.Println("Executing " + c.Name)
}
