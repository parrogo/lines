package main

import (
	"os"

	"github.com/parrogo/lines"
)

func main() {
	var buf lines.Buffer
	for _, arg := range os.Args[1:] {
		buf.AddLine(arg)
		buf.Write(os.Stdout)
	}
}
