package command

import (
	"errors"
)

type CmdDir struct {
	Command
}

func (c CmdDir) CheckParams(Params []string) error {
	var err error

	if err := c.CheckNumberOfParams(Params); err != nil {
		return err
	}

	if err := c.CheckParamValues(Params); err != nil {
		return err
	}

	return err
}

func (c CmdDir) GetName() string {
	return c.Name
}

func (c CmdDir) Execute(Params []string) {
	c.Outputter.Println("Executing " + c.Name)
}

func (c CmdDir) CheckNumberOfParams(Params []string) error {
	var err error

	if len(Params) < 1 {
		err = errors.New("Incorrect syntax.")
	}

	return err
}

func (c CmdDir) CheckParamValues(Params []string) error {
	var err error

	return err
}
