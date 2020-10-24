package session

import (
	"fmt"
	"testing"
)

func TestInitialize(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"asdf", "incorrectRegion"},
		{"eu-central-1", "ok"},
		{"euu-central-1", "incorrectRegion"},
		{"us-east-1", "ok"},
	}

	for _, test := range tests {
		if output, _ := Initialize(test.input); output != test.expected {
			t.Error("Test failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}

	fmt.Println("test session initialization and return type")
}
