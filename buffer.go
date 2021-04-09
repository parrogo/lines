package lines

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
)

// Buffer contains a bytes.Buffer
// and simplify writing lines of text to it.
//
// The empty value of a Buffer is already
// usable.
//
// You can recycle Buffer multiple times,
// any call to WriteFile will truncate the
// underlying bytes.Buffer to 0.
//
// If you want to start from scratch without
// having to write to a file, you can call
// Discard method.
type Buffer struct {
	buf bytes.Buffer
}

// AddLine writes a line of text in the buffer.
// The line to append is built calling `fmt.Sprintf`
// with `lineFormat` and `arguments...`
func (lines *Buffer) AddLine(lineFormat string, arguments ...interface{}) {
	line := fmt.Sprintf(lineFormat, arguments...)
	lines.buf.WriteString(line)
	lines.buf.WriteRune('\n')
}

// WriteFile writes all lines written to
// the buffer so far to a file with given path.
//
// After the file is written successfully, WriteFile
// truncates the internal lines buffer, so that you can recycle
// the same Buffer as if it was newly created.
//
// If an error occurrs during the writing of the file,
// no truncation of the buffer is done, so that you can
// retry the method call later.
func (lines *Buffer) WriteFile(filepath string) error {
	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, fs.FileMode(0644))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(lines.buf.Bytes())
	if err != nil {
		return err
	}

	lines.buf.Truncate(0)

	return err
}

// Write writes all lines written to
// the buffer so far in a io.Writer.
//
// If if the write operation succeed, Write
// truncates the internal lines buffer, so that
// you can recycle the same Buffer as if it
// was newly created.
//
// If an error occurrs during the writing,
// no truncation of the buffer is done, so that you can
// retry the method call later.
func (lines *Buffer) Write(dest io.Writer) error {
	_, err := dest.Write(lines.buf.Bytes())
	if err != nil {
		return err
	}
	lines.buf.Truncate(0)

	return err
}

// Discard truncates the lines buffer, so that you can recycle
// the same Buffer as if it was newly created.
func (lines *Buffer) Discard() {
	lines.buf.Truncate(0)
}
