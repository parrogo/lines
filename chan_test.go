package lines

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteFromChan(t *testing.T) {
	var out bytes.Buffer
	c := make(chan string)
	go func() {
		c <- "line1"
		c <- "line2"
		close(c)
	}()
	assert.NoError(t, WriteFromChan(&out, c))
	assert.Equal(t, "line1\nline2\n", string(out.Bytes()))
}
