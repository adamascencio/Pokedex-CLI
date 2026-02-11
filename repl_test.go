package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " HELLO  adam  ",
			expected: []string{"hello", "adam"},
		},
		{
			input:    "  hello  WORLD its   adam",
			expected: []string{"hello", "world", "its", "adam"},
		},
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of actual doesn't match expected")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word doesn't match expected word")
			}
			fmt.Printf("Test %d: Passed", i)
		}
	}
}
