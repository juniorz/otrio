package otrio

import (
	"reflect"
	"strings"
	"testing"
)

func TestScansOTRMessages(t *testing.T) {
	msgs := []string{
		"hi there",
		"?OTR?",
		"?OTR?v?",
		"?OTRv?",
		"?OTR?v3?",

		//D-H commit
		"?OTR:AAMCLR/vEQAAAAAAAADEw/kAAzQWo5Rm3yf0HGs+hDMvYTo29JOi+75gu92ev1FFjo.",

		// D-H Key
		"?OTR:AAMKN7l3Gi0f7xEAAADAKvkiaOwb4Z+jifIfWqRarnRcRzB+LymYxEyfPxsNp/vjRF.",

		// Reveal sig
		"?OTR:AAMRLR/vETe5dxoAAAAQw55n1Qh6lQ/LM9gMA+BirwAAAdKAplDwy/9tlBeLeDtRxM.",

		// Sig
		"?OTR:AAMSN7l3Gi0f7xEAAAHS8LJL0vXYIoXlLSbItmSYrAJN66S460ZkK6LI7HjzmQBG7u.",

		// Data msg
		"?OTR:AAMDN7l3Gi0f7xEAAAAABAAAAAQAAADAxwZA4PbvhtHOyo+UjzW4JlriUhGG7v1iev.",

		//Error
		"?OTR Error: Something bad happened",

		// Fragments v2
		"?OTR,1,3,?OTR:AAEDAAAAAQAAAAEAAADAVf3Ei72ZgFeKqWvLMnutsQAAAF2SOrDvmZw6g,",
		"?OTR,2,3,JvPUerB9mtf4bqQDFthfoz/XepysnYuReHHEXKe+BFkaEoMNGiBl4TCLZx72gr,",
		"?OTR,3,3,NLKoYOoJTM7zcxsGnvCxaDZCvsmjx3j8Yc5r3i3ylllCQH2/lpr/xCvXFarGtG,",

		// Fragments v3
		"?OTR|c47ba987|00000000,00001,00003,?OTR:AAMCxHuphwAAAAAAAADETecmzCZwU92,",
		"?OTR|c47ba987|00000000,00002,00003,PAfxQZ4+/pUoTPmwj4fLa1bMVyrFBMBJeT7P,",
		"?OTR|c47ba987|00000000,00003,00003,8IOAoR7ZHNHQYPw==.,",
	}

	var b string
	for _, m := range msgs {
		b += m
	}

	s := NewScanner(strings.NewReader(b))

	for _, m := range msgs {
		if !s.Scan() {
			t.Errorf("Failed to scan plain text")
		}

		if r := s.Text(); !reflect.DeepEqual(m, r) {
			t.Errorf("Failed to return message:\n%s\n%s", m, r)
		}
	}
}
