package command

import (
	"errors"
	"strings"
)

type CmdMkDir struct {
	Command
}

func (c CmdMkDir) CheckParams(Params []string) error {
	var err error

	if err := c.CheckNumberOfParams(Params); err != nil {
		return err
	}

	if err := c.CheckParamValues(Params); err != nil {
		return err
	}

	return err
}

func (c CmdMkDir) GetName() string {
	return c.Name
}

func (c CmdMkDir) Execute(Params []string) {
	c.Outputter.Println("Executing " + c.Name)
}

func (c CmdMkDir) CheckNumberOfParams(Params []string) error {
	var err error

	if len(Params) < 1 {
		err = errors.New("Incorrect syntax.")
	}

	return err
}

func (c CmdMkDir) CheckParamValues(Params []string) error {
	var err error

	for i := 0; i < len(Params); i++ {
		err = ParamsContainsBackslashes(Params[i])
		return err
	}

	return err
}

func ParamsContainsBackslashes(Param string) error {
	var err error

	if strings.Contains(Param, "/") || strings.Contains(Param, "\\") {
		err = errors.New("At least one parameter denotes a path rather than a directory name.")
	}

	return err
}
