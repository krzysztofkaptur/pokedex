package main

import "testing"

func TestCleanInput(t *testing.T) {
	type expectedType struct{
		input string
		expected []string
	}

	cases := []expectedType{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "this is Sparta",
			expected: []string{
				"this",
				"is",
				"sparta",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("The lengths are not equal: %v vs %v", len(actual), len(cs.expected))
			continue
		}

		for i := range actual{
			actualWord := actual[i]
			expecedWord := cs.expected[i]

			if actualWord != expecedWord {
				t.Errorf("%v does not equal %v", actualWord, expecedWord)
			}
		}
	}
}