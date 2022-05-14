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
		{"10.9", "11.1", "22"},
		{"10.09", "11.01", "21.1"},
		{"-10.09", "11.01", "0.92"},
		{"11.01", "-10.09", "0.92"},
		{"-11.01", "-10.09", "-21.1"},
		{"-11.01", "10.09", "-0.92"},
		{"0", "6", "6"},
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

func TestLess(t *testing.T) {
	tests := []struct {
		left     string
		right    string
		expected bool
	}{
		{"10", "11", true},
		{"44", "17", false},
		{"44", "171", true},
		{"88", "88", false},
		{"10.01", "111.001", true},
		{"121.1", "10.11", false},
		{"0.1", "0.11", true},
		{"0.14", "0.134", false},
		{"11.1", "12.01", true},
		{"-1", "-2", false},
		{"-1", "2", true},
	}
	for _, tt := range tests {
		left := NewNumber(tt.left)
		right := NewNumber(tt.right)
		result := left.Less(right)
		if result != tt.expected {
			t.Fatalf("Less does not work as expected compering %s and %s less should return %v. got %v", tt.left, tt.right, tt.expected, result)
		}
	}
}

func TestDeduct(t *testing.T) {
	tests := []struct {
		left     string
		right    string
		expected string
	}{
		{"11", "10", "1"},
		{"11.1", "10", "1.1"},
		{"10", "1.1", "8.9"},
		{"10.01", "1.001", "9.009"},
		{"-10.01", "1.001", "-11.011"},
		{"10.01", "-1.001", "11.011"},
		{"-10.01", "-10.001", "-0.009"},
		{"-10.001", "-10.001", "0"},
	}
	for _, tt := range tests {
		left := NewNumber(tt.left)
		right := NewNumber(tt.right)
		result := left.Deduct(right)
		if result.String() != tt.expected {
			t.Fatalf("deducting %s from %s should give result %s. got %s", tt.right, tt.left, tt.expected, result.String())
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		left     string
		right    string
		expected string
	}{
		{"2", "3", "6"},
		{"8", "7", "56"},
		{"11", "10", "110"},
		{"11.1", "10", "111"},
		{"10", "1.1", "11"},
		{"10.01", "1.001", "10.02001"},
		{"-10.01", "1.001", "-10.02001"},
		{"10.01", "-1.001", "-10.02001"},
		{"-10.01", "-10.001", "100.11001"},
		{"-10.001", "-10.001", "100.020001"},
	}
	for _, tt := range tests {
		left := NewNumber(tt.left)
		right := NewNumber(tt.right)
		result := left.Multiply(right)
		if result.String() != tt.expected {
			t.Fatalf("multiplication %s by %s should give result %s. got %s", tt.left, tt.right, tt.expected, result.String())
		}
	}
}

func TestIsNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"123", true},
		{"12g", false},
		{"123456789101112131415161718192021222324252627282930313234353637383940414243456474849", true},
		{"12345678910111213141516171819202122232425262O282930313234353637383940414243456474849", false},
	}
	for _, tt := range tests {
		if isNumber(tt.input) != tt.expected {
			t.Fatalf("isNumber has wrong results: %s should give result %v. got %v", tt.input, isNumber(tt.input), tt.expected)
		}
	}
}

func TestDevide(t *testing.T) {
	tests := []struct {
		left     string
		right    string
		expected string
	}{
		{"6", "3", "2"},
		{"9", "6", "1.5"},
		{"3", "6", "0.5"},
		{"30", "12", "2.5"},
		{"27.072", "12", "2.256"},
		{"176", "4", "44"},
		{"-176", "4", "-44"},
		{"176", "-4", "-44"},
		{"-176", "-4", "44"},
		{"1234567891123456", "2", "617283945561728"},
		{"1234567891123456", "617283945561728", "2"},
		{"12345678911234568", "2", "6172839455617284"},
		{"12345678911234568", "10", "1234567891123456.8"},
		{"1", "3", "0.333333333333333333333333333333"},
		{"1234567891123456888", "288", "4286694066400891.972222222222222222222222222222"},
		{"123456789012345678901234567890.0123456789012345678901234567890", "25556", "4830833816416719318407989.039365015874115715861945927510"},
	}
	for _, tt := range tests {
		left := NewNumber(tt.left)
		right := NewNumber(tt.right)
		result := left.Devide(right)
		if result.String() != tt.expected {
			t.Fatalf("devide %s by %s should give result %s. got %s", tt.left, tt.right, tt.expected, result.String())
		}
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		input     string
		precision uint
		expected  string
	}{
		{"2", 0, "2"},
		{"2.1", 0, "2"},
		{"2.1", 1, "2.1"},
		{"2.11", 1, "2.1"},
		{"2.15", 1, "2.2"},
		{"2.151687", 2, "2.15"},
		{"2.9", 0, "3"},
	}
	for _, tt := range tests {
		input := NewNumber(tt.input)
		result := input.Round(tt.precision)
		if result.String() != tt.expected {
			t.Fatalf("round of  %s by %d should give result %s. got %s", tt.input, tt.precision, tt.expected, result.String())
		}
	}
}

func TestRoundUp(t *testing.T) {
	tests := []struct {
		input     string
		precision uint
		expected  string
	}{
		{"2", 0, "2"},
		{"2.1", 0, "3"},
		{"2.1", 1, "2.1"},
		{"2.11", 1, "2.2"},
		{"2.15", 1, "2.2"},
		{"2.150687", 2, "2.16"},
		{"2.9", 0, "3"},
	}
	for _, tt := range tests {
		input := NewNumber(tt.input)
		result := input.RoundUp(tt.precision)
		if result.String() != tt.expected {
			t.Fatalf("roundUp of  %s by %d should give result %s. got %s", tt.input, tt.precision, tt.expected, result.String())
		}
	}
}

func TestRoundDown(t *testing.T) {
	tests := []struct {
		input     string
		precision uint
		expected  string
	}{
		{"2", 0, "2"},
		{"2.1", 0, "2"},
		{"2.1", 1, "2.1"},
		{"2.11", 1, "2.1"},
		{"2.15", 1, "2.1"},
		{"2.151687", 2, "2.15"},
		{"2.9", 0, "2"},
	}
	for _, tt := range tests {
		input := NewNumber(tt.input)
		result := input.RoundDown(tt.precision)
		if result.String() != tt.expected {
			t.Fatalf("round of  %s by %d should give result %s. got %s", tt.input, tt.precision, tt.expected, result.String())
		}
	}
}

func TestCeil(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"2", "2"},
		{"2.1", "3"},
	}
	for _, tt := range tests {
		input := NewNumber(tt.input)
		result := input.Ceil()
		if result.String() != tt.expected {
			t.Fatalf("ceil of  %s should give result %s. got %s", tt.input, tt.expected, result.String())
		}
	}
}

func TestFloor(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"2", "2"},
		{"2.9", "2"},
	}
	for _, tt := range tests {
		input := NewNumber(tt.input)
		result := input.Floor()
		if result.String() != tt.expected {
			t.Fatalf("floor of  %s should give result %s. got %s", tt.input, tt.expected, result.String())
		}
	}
}
