package command

import (
	"errors"
)

type CmdMkFile struct {
	Command
}

func (c CmdMkFile) CheckParams(Params []string) error {
	var err error

	if err := c.CheckNumberOfParams(Params); err != nil {
		return err
	}

	if err := c.CheckParamValues(Params); err != nil {
		return err
	}

	return err
}

func (c CmdMkFile) GetName() string {
	return c.Name
}

func (c CmdMkFile) Execute(Params []string) {
	c.Outputter.Println("Executing " + c.Name)
}

func (c CmdMkFile) CheckNumberOfParams(Params []string) error {
	var err error

	if len(Params) < 1 {
		err = errors.New("Incorrect syntax.")
	}

	return err
}

func (c CmdMkFile) CheckParamValues(Params []string) error {
	var err error

	return err
}
