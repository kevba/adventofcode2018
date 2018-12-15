package main

import "testing"

func TestReduce(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"dbcCCBcCcD", "dbCBcD"},
		{"daAcCaCAcCcaDA", "daCAcaDA"},
		{"dabAaBAaDA", "daDA"},
		{"abAcCaCBAcCcaA", "abCBAc"},
		{"dabAcCaCBAcCcaDA", "dabCBAcaDA"},
		{"dDabAcCaCBAcCcaDA", "abCBAcaDA"},
		{"dabAcCaCBAcCcaDAa", "dabCBAcaD"},
	}

	for i, test := range tests {
		res := reduce([]byte(test.in))
		if string(res) != test.expected {
			t.Errorf("test %v: result is: %v expected: %v", i, string(res), test.expected)
		}
	}
}
