package number

import "testing"

func TestNumber(t *testing.T) {
	tests := []string{
		"10", "11.1",
	}
	for _, tt := range tests {
		number := NewNumber(tt)
		if number.String() != tt {
			t.Fatalf("number is not well created. expected %s, got %s", tt, number.String())
		}
	}
}

func TestPad(t *testing.T) {
	tests := []struct {
		input     string
		direction int
		length    int
		symbol    string
		expected  string
	}{
		{"10", PADDIRECTIONLEFT, 3, "0", "010"},
		{"10", PADDIRECTIONLEFT, 2, "0", "10"},
		{"10", PADDIRECTIONLEFT, 1, "0", "10"},
		{"10", PADDIRECTIONRIGHT, 3, "0", "100"},
		{"10", PADDIRECTIONRIGHT, 2, "0", "10"},
		{"10", PADDIRECTIONRIGHT, 1, "0", "10"},
	}
	for _, tt := range tests {
		if padded := pad(tt.input, tt.direction, tt.length, tt.symbol); padded != tt.expected {
			t.Fatalf("pad does not work. expected %s. got %s", tt.expected, padded)
		}
	}
}

func TestAddNumber(t *testing.T) {
	tests := []struct {
		left     string
		right    string
		expected string
	}{
		{"10", "11", "21"},
		{"171", "44", "215"},
		{"44", "171", "215"},
		{"88", "88", "176"},
		{"10", "11.1", "21.1"},
		{"10.01", "11.001", "21.011"},
		{"10.9", "11.11", "22.01"},
	}
	for _, tt := range tests {
		left := NewNumber(tt.left)
		right := NewNumber(tt.right)
		result := left.Add(right)
		if result.String() != tt.expected {
			t.Fatalf("adding %s and %s should give result %s. got %s", tt.left, tt.right, tt.expected, result.String())
		}
	}
}
