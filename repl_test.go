package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   Hello World   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   Charmander  Bubasauro  Chalizard",
			expected: []string{"charmander", "bubasauro", "chalizard"},
		},
		{
			input:    " I went camping with my friends!    ",
			expected: []string{"i", "went", "camping", "with", "my", "friends!"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(c.expected) != len(actual) {
			t.Errorf("Actual size %v is diferent from expected %v", len(actual), len(c.expected))
		}
		for i, word := range actual {
			if word != c.expected[i] {
				t.Errorf("Word %v is difernt from expected word %v", word, c.expected[i])
			}
		}
	}
}
