package main

import "testing"

func TestReduce(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"a", "a"},
		{"aA", ""},
		{"aa", "aa"},
		{"aaa", "aaa"},
		{"Aa", ""},
		{"AaAaAaAaAaAaAa", ""},
		{"abba", "abba"},
		{"aBba", "aa"},
		{"aBbBa", "aBa"},
		{"ABba", ""},
		{"cABbaC", ""},
	}

	for i, test := range tests {
		res := reduce([]byte(test.in))
		if string(res) != test.expected {
			t.Errorf("test %v: result is: %v expected: %v", i, string(res), test.expected)
		}
	}
}
