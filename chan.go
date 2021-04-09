package lines

import (
	"io"
)

// WriteFromChan continuously writes
// to w everything received from c chan.
func WriteFromChan(w io.Writer, c chan string) error {
	var buf Buffer
	for line := range c {
		buf.AddLine(line)
		_, err := w.Write(buf.Reset())
		if err != nil {
			return err
		}
	}
	return nil
}
