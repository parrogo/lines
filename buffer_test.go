package lines

import (
	"io/fs"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinesBuf(t *testing.T) {
	content, err := fs.ReadFile(fixtureFS, "placeholder")
	var buf Buffer
	if assert.NoError(t, err) {
		for _, l := range strings.Split(string(content), "\n") {
			buf.AddLine(l)
		}
	}

	err = os.Remove("/tmp/testlines")
	if !os.IsNotExist(err) && assert.NoError(t, err) {
		return
	}
	err = buf.WriteFile("/tmp/testlines")
	if assert.NoError(t, err) {
		return
	}

	actual, err := fs.ReadFile(fixtureFS, "/tmp/testlines")
	if assert.NoError(t, err) {
		return
	}

	assert.Equal(t, actual, content)

	err = os.Remove("/tmp/testlines")
	if assert.NoError(t, err) {
		return
	}

}
