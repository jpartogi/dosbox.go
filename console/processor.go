package console

import (
	"bufio"
	"fmt"
	"github.com/jpartogi/dosboxgo/filesystem"
	"os"
	"strings"
)

type Processor struct {
	Drive     filesystem.Drive
	Outputter Outputter
}

func NewProc(Drive filesystem.Drive, Outputter Outputter) Processor {
	return Processor{Drive, Outputter}
}

func (p *Processor) ProcessInput() {
	p.Outputter.Println("DOSBox, Scrum.org, Professional Scrum Developer Training")
	p.Outputter.Println("Copyright (c) Joshua Partogi. All rights reserved.")

	line := ""
	reader := bufio.NewReader(os.Stdin)

	for strings.ToLower(line) != "exit" {
		fmt.Print("C:\\>")
		text, _ := reader.ReadString('\n')
		line = strings.TrimSpace(text)
		fmt.Println(line)
	}
}
