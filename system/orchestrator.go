package system

import "fmt"
import "dosbox/filesystem"

func OrchestrateSystem() {
	fmt.Println("Hello")
	filesystem.Restore()
}
