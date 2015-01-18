package tests

import (
	"github.com/jpartogi/dosboxgo/command"
	"testing"
)

func TestParamsContainsBackslashes(t *testing.T) {
	var err error

	err = command.ParamsContainsBackslashes("baz foo")
	if err != nil {
		t.Error("Error not supposed to be thrown.")
	}

	err = command.ParamsContainsBackslashes("baz /")
	if err == nil {
		t.Error("Error is not thrown.")
	}

	err = command.ParamsContainsBackslashes("baz \\")
	if err == nil {
		t.Error("Error is not thrown.")
	}

}
