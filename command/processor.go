package command

import (
	"bufio"
	"fmt"
	"github.com/jpartogi/dosboxgo/console"
	"github.com/jpartogi/dosboxgo/filesystem"
	"os"
	"strings"
)

type Processor struct {
	Invoker   Invoker
	Drive     filesystem.Drive
	Outputter console.Outputter
}

func NewProc(Invoker Invoker, Drive filesystem.Drive, Outputter console.Outputter) Processor {
	return Processor{Invoker, Drive, Outputter}
}

func (p *Processor) ProcessInput() {
	p.Outputter.Println("DOSBox, Scrum.org, Professional Scrum Developer Training")
	p.Outputter.Println("Copyright (c) Joshua Partogi. All rights reserved.")

	line := ""
	reader := bufio.NewReader(os.Stdin)

	for strings.ToLower(line) != "exit" {
		p.Outputter.NewLine()
		fmt.Print("C:\\>")

		text, _ := reader.ReadString('\n')
		line = strings.TrimSpace(text)

		p.Invoker.ExecuteCommand(line)
	}
}
