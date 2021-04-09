package lines_test

import (
	"embed"
	"fmt"
	"io/fs"
	
	"github.com/parrogo/lines"
)

//go:embed fixtures
var fixtureRootFS embed.FS
var fixtureFS, _ = fs.Sub(fixtureRootFS, "fixtures")

// This example show how to use lines.Func()
func ExampleFunc() {
	fmt.Println(lines.Func())
	// Output: 42
}
