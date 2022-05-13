package number

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	COMPAREEQUAL = iota
	COMPARELESS
	COMPAREMORE
)

const (
	PADDIRECTIONRIGHT = iota
	PADDIRECTIONLEFT
)

func NewNumber(pNumber string) Number {
	number := pNumber
	result := Number{}
	if number == "" {
		number = "0"
	}
	if number[0] == '-' {
		number = string(number[1:])
		result.negative = true
	}
	parts := strings.Split(number, ".")
	if len(parts) > 2 {
		log.Fatalf("number conversion error, gor %d parts of the number %s", len(parts), number)
	}
	_, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatalf("error while converting string %s to integer: %s", parts[0], err)
	}
	result.whole = parts[0]
	if len(parts) > 1 {
		_, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("error while converting string %s to integer: %s", parts[1], err)
		}
		result.reminder = parts[1]
	}
	return result
}

type Number struct {
	negative bool
	whole    string
	reminder string
}

func (n Number) String() string {
	if n.reminder == "" {
		return n.whole
	}
	var sign string
	if n.negative {
		sign = "-"
	}
	return fmt.Sprintf("%s%s.%s", sign, n.whole, n.reminder)
}

func pad(s string, direction int, length int, symbol string) string {
	result := s

	for i := 0; i < length-len(s); i++ {
		switch direction {
		case PADDIRECTIONLEFT:
			result = fmt.Sprintf("%s%s", symbol, result)
		case PADDIRECTIONRIGHT:
			result = fmt.Sprintf("%s%s", result, symbol)
		}
	}

	return result
}

func sumString(pLeft, pRight string, pMemory int) (result string, memory int) {
	memory = pMemory
	for i := len(pLeft) - 1; i >= 0; i-- {
		x, _ := strconv.Atoi(string(pLeft[i]))
		y, _ := strconv.Atoi(string(pRight[i]))
		sum := strconv.Itoa(x + y + memory)
		memory = 0
		var digit string
		if len(sum) > 1 {
			memory, _ = strconv.Atoi(string(sum[0]))
			digit = string(sum[1])
		} else {
			digit = sum
		}
		result = fmt.Sprintf("%s%s", digit, result)
	}

	return result, memory
}

func addReminder(pLeft, pRight string) (string, int) {
	left := pad(pLeft, PADDIRECTIONRIGHT, len(pRight), "0")
	right := pad(pRight, PADDIRECTIONRIGHT, len(pLeft), "0")

	result, reminder := sumString(left, right, 0)

	for result[len(result)-1] == '0' {
		result = result[:len(result)-1]
		if len(result) == 0 {
			break
		}
	}
	return result, reminder
}

func addWhole(pLeft, pRight string, pMemory int) (result string) {
	left := pad(pLeft, PADDIRECTIONLEFT, len(pRight), "0")
	right := pad(pRight, PADDIRECTIONLEFT, len(pLeft), "0")

	result, memory := sumString(left, right, pMemory)

	if memory > 0 {
		digit := strconv.Itoa(memory)
		result = fmt.Sprintf("%s%s", digit, result)
	}

	return result
}

func (n *Number) Abs() Number {
	result := NewNumber("0")
	result.whole = n.whole
	result.reminder = n.reminder
	return result
}

func (n *Number) Add(a Number) Number {
	result := NewNumber("")
	if n.negative == a.negative {
		var memory int = 0
		isReminder := n.reminder != "" || a.reminder != ""
		if isReminder {
			result.reminder, memory = addReminder(n.reminder, a.reminder)
		}
		result.whole = addWhole(n.whole, a.whole, memory)
		result.negative = n.negative
	} else {
		if n.negative {
			result = a.Deduct(n.Abs())
		} else {
			result = n.Deduct(a.Abs())
		}
	}
	return result
}

func deductString(pLeft, pRight string, pMemory int) (result string, memory int) {
	memory = pMemory
	for i := len(pLeft) - 1; i >= 0; i-- {
		x, _ := strconv.Atoi(string(pLeft[i]))
		y, _ := strconv.Atoi(string(pRight[i]))
		amt := x - y + memory
		memory = 0
		if amt < 0 {
			amt += 10
			memory = -1
		}
		digit := strconv.Itoa(amt)
		result = fmt.Sprintf("%s%s", digit, result)
	}

	return result, memory
}

func deductReminder(pLeft, pRight string) (string, int) {
	left := pad(pLeft, PADDIRECTIONRIGHT, len(pRight), "0")
	right := pad(pRight, PADDIRECTIONRIGHT, len(pLeft), "0")

	result, reminder := deductString(left, right, 0)

	for result[len(result)-1] == '0' {
		result = result[:len(result)-1]
		if len(result) == 0 {
			break
		}
	}
	return result, reminder
}

func deductWhole(pLeft, pRight string, pMemory int) (result string) {
	left := pad(pLeft, PADDIRECTIONLEFT, len(pRight), "0")
	right := pad(pRight, PADDIRECTIONLEFT, len(pLeft), "0")

	result, _ = deductString(left, right, pMemory)
	for result[0] == '0' && len(result) > 1 {
		result = result[1:]
	}

	return result
}

func (n *Number) Copy() Number {
	result := NewNumber("0")
	result.negative = n.negative
	result.whole = n.whole
	result.reminder = n.reminder
	return result
}

func (n *Number) Deduct(a Number) Number {
	if n.negative && !a.negative {
		negative := NewNumber(fmt.Sprintf("-%s", a.String()))
		return n.Add(negative)
	} else if !n.negative && a.negative {
		return n.Add(a.Abs())
	}
	isFirstSMaller := n.uLess(a)
	left := n.Copy()
	right := a.Copy()
	if isFirstSMaller {
		left = right
		left.negative = !left.negative
		right = n.Copy()
	}
	result := NewNumber("")
	var memory int = 0
	isReminder := left.reminder != "" || right.reminder != ""
	if isReminder {
		result.reminder, memory = deductReminder(left.reminder, right.reminder)
	}
	result.whole = deductWhole(left.whole, right.whole, memory)
	result.negative = left.negative

	return result
}

func stringCompare(left, right string) int {
	for i := 0; i < len(left); i++ {
		leftDigit, _ := strconv.Atoi(string(left[i]))
		rightDigit, _ := strconv.Atoi(string(right[i]))
		if leftDigit > rightDigit {
			return COMPAREMORE
		} else if rightDigit > leftDigit {
			return COMPARELESS
		}
	}
	return COMPAREEQUAL
}

func reminderCompare(pLeft, pRight string) int {
	left := pad(pLeft, PADDIRECTIONRIGHT, len(pRight), "0")
	right := pad(pRight, PADDIRECTIONRIGHT, len(pLeft), "0")
	return stringCompare(left, right)
}

func wholeCompare(pLeft, pRight string) int {
	left := pad(pLeft, PADDIRECTIONLEFT, len(pRight), "0")
	right := pad(pRight, PADDIRECTIONLEFT, len(pLeft), "0")
	return stringCompare(left, right)
}

func (n *Number) uLess(a Number) bool {
	switch wholeCompare(n.whole, a.whole) {
	case COMPAREMORE:
		return false
	case COMPARELESS:
		return true
	case COMPAREEQUAL:
		switch reminderCompare(n.reminder, a.reminder) {
		case COMPARELESS:
			return true
		default:
			return false
		}
	}
	return false
}

func (n *Number) Less(a Number) bool {
	if n.negative && a.negative {
		return !n.uLess(a)
	} else if !n.negative && !a.negative {
		return n.uLess(a)
	} else if !a.negative {
		return true
	}
	return false
}

func (n *Number) More(a Number) bool {
	return !n.Less(a)
}
