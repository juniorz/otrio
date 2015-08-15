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

	if len(data) <= len(otrMarker) {
		return 0, nil, nil
	}

	p := 0
	marker := bytes.Index(data, otrMarker)

	switch marker {
	case -1:
		p = len(data)
	case 0:
		//TODO: handle atEOF inside every function when it does not advance the scanner position
		p = len(otrMarker)
		switch data[p] {
		case '|', ',':
			return scanFragment(data, atEOF)
		case ':':
			return scanEncodedMessage(data, atEOF)
		case '?', 'v':
			return scanQueryMessageBody(data, atEOF)
		case ' ':
			//Error messages
			//Its hard to know when they have ended if they have details
			//and are followed by a plain message
			//We consider the message only ends at the next OTR marker
			i := bytes.Index(data[p:], otrMarker)
			if i == -1 {
				//TODO: consider atEOF
				return 0, nil, nil
			}

			return p + i, data[:p+i], nil
		}
	default:
		p = marker
	}

	return p, data[:p], nil
}

func scanFragment(data []byte, atEOF bool) (advance int, token []byte, err error) {
	p := len(otrMarker)
	for j := 0; j < 4; j++ {
		i := bytes.IndexByte(data[p:], ',') + 1
		if i == 0 {
			if atEOF {
				return p, data[:p], io.EOF
			}

			return 0, nil, nil
		}

		p += i
	}

	return p, data[:p], nil
}

func scanEncodedMessage(data []byte, atEOF bool) (advance int, token []byte, err error) {
	p := len(otrMarker)
	i := bytes.IndexByte(data[p:], '.') + 1
	if i == 0 {
		return 0, nil, nil
	}

	return p + i, data[:p+i], nil
}

func scanQueryMessageBody(data []byte, atEOF bool) (advance int, token []byte, err error) {
	p := len(otrMarker)
	if data[p] == '?' {
		p++
	}

	if len(data) < p {
		return 0, nil, nil
	}

	if data[p] == 'v' {
		p++

		i := bytes.IndexByte(data[p:], '?')
		if i == -1 {
			return 0, nil, nil
		}

		p += i + 1
	}

	return p, data[:p], nil
}
