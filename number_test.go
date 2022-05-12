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
