package console

import "fmt"

type Outputter struct {
	NumberOfPrintedChars int
}

func NewOutputter() Outputter {
	return Outputter{0}
}

func (o *Outputter) Println(line string) {
	fmt.Println(line)
}

func (o *Outputter) Print(line string) {
	fmt.Print(line)
}
