package tests

import (
	"github.com/jpartogi/dosboxgo/command"
	"testing"
)

func TestParseCommand(t *testing.T) {
	commandName, commandParams := command.ParseCommand("mkdir bar foo baz")

	if commandName != "mkdir" {
		t.Error("mkdir command expected")
	}

	if commandParams[0] != "bar" {
		t.Error("bar param expected")
	}

	if commandParams[1] != "foo" {
		t.Error("foo param expected")
	}

	if commandParams[2] != "baz" {
		t.Error("baz param expected")
	}
}
