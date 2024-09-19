package bom_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/spkg/bom"
)

var testCases = []struct {
	Input    []byte
	Expected []byte
}{
	{
		Input:    nil,
		Expected: nil,
	},
	{
		Input:    []byte{},
		Expected: []byte{},
	},
	{
		Input:    []byte{0xef},
		Expected: []byte{0xef},
	},
	{
		Input:    []byte{0xef, 0xbb},
		Expected: []byte{0xef, 0xbb},
	},
	{
		Input:    []byte{0xef, 0xbb, 0xbf},
		Expected: []byte{},
	},
	{
		Input:    []byte{0xef, 0xbb, 0xbf, 0x41, 0x42, 0x43},
		Expected: []byte{0x41, 0x42, 0x43},
	},
	{
		Input:    []byte{0xef, 0xbb, 0x41, 0x42, 0x43},
		Expected: []byte{0xef, 0xbb, 0x41, 0x42, 0x43},
	},
	{
		Input:    []byte{0xef, 0x41, 0x42, 0x43},
		Expected: []byte{0xef, 0x41, 0x42, 0x43},
	},
	{
		Input:    []byte{0x41, 0x42, 0x43},
		Expected: []byte{0x41, 0x42, 0x43},
	},
}

func TestClean(t *testing.T) {
	for _, tc := range testCases {
		actual := bom.Clean(tc.Input)
		if !bytes.Equal(actual, tc.Expected) {
			t.Errorf("got %v, want %v", actual, tc.Expected)
		}
	}
}

func TestReader(t *testing.T) {
	for _, tc := range testCases {
		// An input value of nil works differently to the Clean function.
		// In this case it results in an empty buffer, not nil.
		expected := tc.Expected
		if tc.Input == nil {
			expected = []byte{}
		}
		r1 := bytes.NewReader(tc.Input)
		r2 := bom.NewReader(r1)
		actual, err := io.ReadAll(r2)
		if err != nil {
			t.Errorf("want no error, got %v", err)
			continue
		}
		if !bytes.Equal(actual, expected) {
			t.Errorf("got %v, want %v", actual, expected)
		}
	}
}
