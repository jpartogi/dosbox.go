package tests

import (
	"github.com/jpartogi/dosboxgo/filesystem"
	"testing"
)

func TestCurrentDirectory(t *testing.T) {
	drive := filesystem.NewDrive("C")
	drive.Restore()
}
