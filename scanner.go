package otrio

import (
	"bufio"
	"bytes"
	"io"
)

var (
	otrMarker = []byte("?OTR")
)

// NewScanner returns a new Scanner to read from r.
// The split function defaults to ScanOTR
func NewScanner(r io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(r)
	scanner.Split(ScanOTR)
	return scanner
}

// ScanOTR is a split function for a bufio.Scanner that returns each OTR message
// of a text. Strings in between OTR messages are considered to be plain-text
// messages, and are returned as soon as the enclosing OTR message is encountered
// or the EOF is reached.
func ScanOTR(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return len(data), data, io.EOF
	}

	p := 0
	marker := bytes.Index(data, otrMarker)

	switch marker {
	case -1:
		p = len(data)
	case 0:
		// might need more data
		// in this case, should advance 0 bytes

		p = len(otrMarker)

		switch data[p] {
		case '|', ',':
			for j := 0; j < 4; j++ {
				i := bytes.IndexByte(data[p:], ',')
				if i == -1 {
					if atEOF {
						return p, data[:p], io.EOF
					}

					return 0, nil, nil
				}

				p += i + 1
			}
		case ':':
			i := bytes.IndexByte(data[p:], '.')
			if i == -1 {
				return 0, nil, nil
			}

			p += i + 1
		case '?', 'v':
			if data[p] == '?' {
				p++
			}

			if data[p] == 'v' {
				p++

				i := bytes.IndexByte(data[p:], '?')
				if i == -1 {
					return 0, nil, nil
				}

				p += i + 1
			}
		}
	default:
		p = marker
	}

	token = data[:p]
	advance = p

	return
}
