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
		{"aABbcCdDaADAad", ""},
	}

	for i, test := range tests {
		res := reduce([]byte(test.in))
		if string(res) != test.expected {
			t.Errorf("test %v: result is: %v expected: %v", i, string(res), test.expected)
		}
	}
}

func TestStripBytes(t *testing.T) {
	tests := []struct {
		in       string
		remove   string
		expected string
	}{
		{"dabAcCaCBAcCcaDA", "a", "dbcCCBcCcD"},
		{"dabAcCaCBAcCcaDA", "b", "daAcCaCAcCcaDA"},
		{"dabAcCaCBAcCcaDA", "c", "dabAaBAaDA"},
		{"dabAcCaCBAcCcaDA", "A", "dbcCCBcCcD"},
		{"dabAcCaCBAcCcaDA", "B", "daAcCaCAcCcaDA"},
		{"dabAcCaCBAcCcaDA", "C", "dabAaBAaDA"},
	}

	for i, test := range tests {
		res := stripBytes([]byte(test.in), byte(test.remove[0]))
		if string(res) != test.expected {
			t.Errorf("test %v: result is: %v expected: %v", i, string(res), test.expected)
		}
	}
}

func TestStripAndReduce(t *testing.T) {
	tests := []struct {
		in       string
		remove   string
		expected string
	}{
		{"dabAcCaCBAcCcaDA", "a", "dbCBcD"},
		{"dabAcCaCBAcCcaDA", "b", "daCAcaDA"},
		{"dabAcCaCBAcCcaDA", "c", "daDA"},
		{"dabAcCaCBAcCcaDA", "d", "abCBAc"},
		{"dabAcCaCBAcCcaDA", "A", "dbCBcD"},
		{"dabAcCaCBAcCcaDA", "B", "daCAcaDA"},
		{"dabAcCaCBAcCcaDA", "C", "daDA"},
		{"dabAcCaCBAcCcaDA", "D", "abCBAc"},
	}

	for i, test := range tests {
		res := reduce(stripBytes([]byte(test.in), byte(test.remove[0])))
		if string(res) != test.expected {
			t.Errorf("test %v: result is: %v expected: %v", i, string(res), test.expected)
		}
	}

}
