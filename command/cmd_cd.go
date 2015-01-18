package command

import (
	"errors"
)

type CmdCd struct {
	Command
}

func (c CmdCd) CheckParams(Params []string) error {
	var err error

	if err := c.CheckNumberOfParams(Params); err != nil {
		return err
	}

	if err := c.CheckParamValues(Params); err != nil {
		return err
	}

	return err
}

func (c CmdCd) GetName() string {
	return c.Name
}

func (c CmdCd) Execute(Params []string) {
	c.Outputter.Println("Executing " + c.Name)
}

func (c CmdCd) CheckNumberOfParams(Params []string) error {
	var err error

	if len(Params) < 1 {
		err = errors.New("Incorrect syntax.")
	}

	return err
}

func (c CmdCd) CheckParamValues(Params []string) error {
	var err error

	return err
}
