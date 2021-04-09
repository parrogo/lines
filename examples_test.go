package lines_test

import (
	"embed"
	"io/fs"
	"os"

	"github.com/parrogo/lines"
)

//go:embed fixtures
var fixtureRootFS embed.FS
var fixtureFS, _ = fs.Sub(fixtureRootFS, "fixtures")

// This example show how to use a lines.Buffer
func ExampleBuffer() {
	var lines lines.Buffer
	lines.AddLine("A simple line")
	lines.AddLine("Trailing new lines are not needed\n")
	lines.AddLine("You can use %s arguments", "fmt.Printf")
	lines.Write(os.Stdout)
	// Output: A simple line
	// Trailing new lines are not needed
	//
	// You can use fmt.Printf arguments
}
