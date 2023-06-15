package lexorank

import (
	"testing"
)

func Test_incrementChar(t *testing.T) {
	testCases := []struct {
		input  byte
		output byte
		ok     bool
	}{
		{'0', '1', true},
		{'9', 'A', true},
		{'Z', 'a', true},
		{'z', 0, false},
	}

	for _, tc := range testCases {
		output, ok := incrementChar(tc.input)
		if output != tc.output || ok != tc.ok {
			t.Errorf("incrementChar(%c) => (%c, %v), want (%c, %v)", tc.input, output, ok, tc.output, tc.ok)
		}
	}
}

func Test_decrementChar(t *testing.T) {
	testCases := []struct {
		input  byte
		output byte
		ok     bool
	}{
		{'1', '0', true},
		{'A', '9', true},
		{'a', 'Z', true},
		{'0', 0, false},
	}

	for _, tc := range testCases {
		output, ok := decrementChar(tc.input)
		if output != tc.output || ok != tc.ok {
			t.Errorf("decrementChar(%c) => (%c, %v), want (%c, %v)", tc.input, output, ok, tc.output, tc.ok)
		}
	}
}
