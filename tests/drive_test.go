package tests

import "testing"
import "github.com/jpartogi/dosboxgo/filesystem"

func TestCurrentDirectory(t *testing.T) {
	drive := filesystem.NewDrive("C")
	drive.Restore()
}
