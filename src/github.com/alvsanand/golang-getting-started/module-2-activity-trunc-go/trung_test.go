package main

import (
	"testing"
)
func TestGetIntegerFromFloatingNumber(t *testing.T) {
    var tests = []struct {
		input    string
		expected (int64)
	}{
		{"", 0 },
		{"123", 123 },
		{"123.45", 123 },
		{".45", 0 },
	}

	for _, test := range tests {
		if output, _ := GetIntegerFromFloatingNumber(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}