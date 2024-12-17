package calculator

import (
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
		expectErr  bool
	}{
		{"2 + 2", 4, false},
		{"3 * 3", 9, false},
		{"4 / 2", 2, false},
		{"10 - 5+5", 10, false},
		{"5 / 0 * 4", 0, true},
		{"invalid", 0, true},
	}

	for _, tt := range tests {
		result, err := Calc(tt.expression)
		if (err != nil) != tt.expectErr {
			t.Errorf("Calc(%q) expected error: %v, got: %v", tt.expression, tt.expectErr, err)
		}
		if !tt.expectErr && result != tt.expected {
			t.Errorf("Calc(%q) = %v; expected %v", tt.expression, result, tt.expected)
		}
	}
}
