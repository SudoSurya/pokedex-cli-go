package main

import (
	"testing"
)

func TestCleanInpur(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		}, {
			input: "helLo world",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, curr := range cases {
		actual := cleanInput(curr.input)
		if len(actual) != len(curr.expected) {
			t.Errorf(
				"the lengths isn't macthing %v %v",
				len(curr.expected),
				len(curr.input),
			)
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := curr.expected[i]

			if actualWord != expectedWord {
				t.Errorf("%v does not equal %v",
					actualWord,
					expectedWord,
				)
			}
		}
	}
}
