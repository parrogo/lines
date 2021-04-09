package lines_test

import (
	"embed"
	"fmt"
	"io/fs"
	"os"

	"github.com/parrogo/lines"
)

//go:embed fixtures
var fixtureRootFS embed.FS
var fixtureFS, _ = fs.Sub(fixtureRootFS, "fixtures")

// This example show how to use a lines.Buffer
// to acccumulate a bunch of lines and write them
// to console at end using `Write` method.
func ExampleBuffer_Write() {
	var lines lines.Buffer
	lines.AddLineF("A simple line")
	lines.AddLineF("Trailing new lines are not needed\n")
	lines.AddLineF("You can use %s arguments", "fmt.Printf")

	err := lines.Write(os.Stdout)
	if err != nil {
		panic(err)
	}
	// Output: A simple line
	// Trailing new lines are not needed
	//
	// You can use fmt.Printf arguments
}

// This example show how to use a lines.Buffer
func ExampleBuffer_WriteFile() {
	var lines lines.Buffer
	lines.AddLineF("A simple line")
	lines.AddLineF("Trailing new lines are not needed\n")
	lines.AddLineF("You can use %s arguments", "fmt.Printf")

	err := lines.WriteFile("/tmp/example")
	if err != nil {
		panic(err)
	}
	// read the content of the file just written
	content, _ := os.ReadFile("/tmp/example")
	os.Remove("/tmp/example")
	fmt.Print(string(content))
	// Output: A simple line
	// Trailing new lines are not needed
	//
	// You can use fmt.Printf arguments
}
