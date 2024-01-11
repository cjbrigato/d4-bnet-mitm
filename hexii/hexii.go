package hexii

import (
	"bytes"
	"fmt"
	"io"
)

// Dump returns a HexII formatted string representing the given data. For
// more control over the output, use a Dumper.
func Dump(data []byte) string {
	buf := new(bytes.Buffer)
	d := Dumper(buf, 16, 4, false)
	d.Write(data)
	d.Close()
	return buf.String()
}

// Dumper returns a new io.WriteCloser that streams the given data to the
// output in HexII format. The width parameter set the number of bytes shown
// per line, offsetDigits sets how many digits are used when formatting
// offsets, and the header bool controls whether a ruler is printed above
// the hex dump.
func Dumper(out io.Writer, width, offsetDigits int, header bool) io.WriteCloser {
	if width < 1 {
		panic("width must be >= 1")
	}

	d := &dumper{
		out:       out,
		width:     width,
		header:    header,
		lineLen:   3 * width,
		cleanLine: make([]byte, 3*width),
		buffer:    new(bytes.Buffer),
		readbuf:   make([]byte, width),
		prefixLen: offsetDigits,
	}
	for i := range d.cleanLine {
		d.cleanLine[i] = ' '
	}
	return d
}

type dumper struct {
	out       io.Writer
	width     int
	header    bool
	lineLen   int
	cleanLine []byte
	buffer    *bytes.Buffer
	readbuf   []byte
	prefixLen int

	offset  int
	line    []byte
	repr    []byte
	oldPref []byte
}

func (d *dumper) Write(data []byte) (int, error) {
	d.buffer.Write(data)
	d.process()
	return len(data), nil
}

func (d *dumper) Close() error {
	d.final()
	return nil
}

func (d *dumper) process() {
	for d.buffer.Len() >= d.width {
		if d.header && d.offset == 0 {
			d.ruler()
		}

		d.buffer.Read(d.readbuf)

		d.line = appendLine(d.line[:0], d.readbuf)
		if bytes.Equal(d.line, d.cleanLine) {
			d.oldPref = nil
			d.offset += d.width
			continue
		}

		pref := []byte(fmt.Sprintf("%0*x:", d.prefixLen, d.offset))
		d.repr = append(d.repr[:0], cleanPrefix(pref, d.oldPref)...)
		d.repr = append(d.repr, d.line...)
		d.repr = append(d.repr, '\n')
		d.out.Write(d.repr)
		d.oldPref = pref
		d.offset += d.width
	}
}

func (d *dumper) final() {
	if d.header && d.offset == 0 {
		d.ruler()
	}

	n, _ := d.buffer.Read(d.readbuf)

	d.line = appendLine(d.line[:0], d.readbuf[:n])
	pref := []byte(fmt.Sprintf("%0*x:", d.prefixLen, d.offset))
	d.repr = append(d.repr[:0], cleanPrefix(pref, d.oldPref)...)
	d.repr = append(d.repr, d.line...)
	d.repr = append(d.repr, ' ', ']', '\n')
	d.out.Write(d.repr)
}

func (d *dumper) ruler() {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "% *s ", d.prefixLen, "")
	for i := 0; i < d.width; i++ {
		fmt.Fprintf(buf, " %2x", i)
	}
	fmt.Fprintf(buf, "\n\n")
	buf.WriteTo(d.out)
}

func cleanPrefix(cur, prev []byte) []byte {
	c := make([]byte, len(cur))
	copy(c, cur)
	for i := 0; i < len(prev) && cur[i] == prev[i]; i++ {
		c[i] = ' '
	}
	return c
}

func appendLine(bs, data []byte) []byte {
	for _, c := range data {
		bs = append(bs, ' ')
		bs = appendRepr(bs, c)
	}
	return bs
}

func appendRepr(bs []byte, char byte) []byte {
	if char == 0 {
		return append(bs, ' ', ' ')
	}
	if char < 32 {
		return appendNumeric(bs, char)
	}
	if char < 127 {
		return append(bs, '.', char)
	}
	if char == 255 {
		return append(bs, '#', '#')
	}
	return appendNumeric(bs, char)
}

const digits = "0123456789abcdef"

func appendNumeric(bs []byte, char byte) []byte {
	return append(bs, digits[char>>4], digits[char&0xf])
}
