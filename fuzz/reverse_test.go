package main

import (
	"testing"
	"unicode/utf8"
)

// unit test of Reverse
func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}

	for _, tc := range testcases {
		rev, revErr := Reverse(tc.in)
		if revErr != nil {
			return
		}
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

// fuzz test begins with FuzzXxxx instead of TestXxx(unittest)
// takes *testing.F instead of *testing.T
// can run the fuzz test by go test -run=FuzzReverse if there are more than 1 test
func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // use f.Add to provide a see corpus
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
