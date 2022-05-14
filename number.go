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

func isNumber(pInput string) bool {
	const BIGGESTINTLENGTH int = 18
	if len(pInput) > BIGGESTINTLENGTH {
		input := pInput
		for len(input) > 0 {
			var reminder string
			if len(input) >= BIGGESTINTLENGTH {
				reminder = input[BIGGESTINTLENGTH:]
				input = input[:BIGGESTINTLENGTH]
			}
			if !isNumber(input) {
				return false
			}
			input = reminder
		}
	} else {
		_, err := strconv.Atoi(pInput)
		if err != nil {
			return false
		}
	}
	return true
}

func NewNumber(pNumber string) Number {
	number := pNumber
	result := Number{}
	if pNumber != "" {
		if number[0] == '-' {
			number = string(number[1:])
			result.negative = true
		}
		parts := strings.Split(number, ".")
		if len(parts) > 2 {
			log.Fatalf("number conversion error, got %d parts of the number %s", len(parts), number)
		}
		if !isNumber(parts[0]) {
			log.Fatalf("string %s is not a number", parts[0])
		}
		result.whole = parts[0]
		if len(parts) > 1 {
			if !isNumber(parts[1]) {
				log.Fatalf("string %s is not a number", parts[1])
			}
			result.reminder = parts[1]
		}
	}
	return result
}

type Number struct {
	negative bool
	whole    string
	reminder string
}

func (n Number) String() string {
	whole := n.whole
	if whole == "" {
		whole = "0"
	}
	var sign string
	if n.negative {
		sign = "-"
	}
	if n.reminder == "" {
		return fmt.Sprintf("%s%s", sign, whole)
	}
	return fmt.Sprintf("%s%s.%s", sign, whole, n.reminder)
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
	if result.String() == "-0" {
		result.negative = false
	}

	return result
}

func (n *Number) Multiply(a Number) Number {
	left := fmt.Sprintf("%s%s", n.whole, n.reminder)
	right := fmt.Sprintf("%s%s", a.whole, a.reminder)
	decimals := len(n.reminder) + len(a.reminder)
	result := NewNumber("0")
	decimalIndex := 0
	for i := len(right) - 1; i >= 0; i-- {
		multi := pad("", PADDIRECTIONLEFT, decimalIndex, "0")
		var memory int
		for j := len(left) - 1; j >= 0; j-- {
			leftDigit, _ := strconv.Atoi(string(left[j]))
			rightDigit, _ := strconv.Atoi(string(right[i]))
			amt := strconv.Itoa(leftDigit*rightDigit + memory)
			memory = 0
			if len(amt) > 1 {
				memory, _ = strconv.Atoi(string(amt[0]))
				multi = fmt.Sprintf("%s%s", string(amt[1]), multi)
			} else {
				multi = fmt.Sprintf("%s%s", string(amt[0]), multi)
			}

		}
		if memory > 0 {
			multi = fmt.Sprintf("%d%s", memory, multi)
			memory = 0
		}
		result = result.Add(NewNumber(multi))
		decimalIndex++
	}
	result.reminder = result.whole[len(result.whole)-decimals:]
	result.whole = result.whole[:len(result.whole)-decimals]
	if len(result.reminder) > 0 {
		for result.reminder[len(result.reminder)-1] == '0' {
			result.reminder = result.reminder[:len(result.reminder)-1]
			if len(result.reminder) == 0 {
				break
			}
		}
	}
	if n.negative != a.negative {
		result.negative = true
	}

	return result
}

func (n *Number) Devide(a Number) Number {
	reminderLength := len(n.reminder)
	if len(a.reminder) > reminderLength {
		reminderLength = len(a.reminder)
	}
	left := NewNumber(fmt.Sprintf("%s%s", n.whole, pad(n.reminder, PADDIRECTIONRIGHT, reminderLength, "0")))
	right := NewNumber(fmt.Sprintf("%s%s", a.whole, pad(a.reminder, PADDIRECTIONRIGHT, reminderLength, "0")))

	result := NewNumber("")
	if n.negative != a.negative {
		result.negative = true
	}

	var digit int
	var isReminder bool
	var leftover string
	if len(left.whole) > len(right.whole) {
		leftover = left.whole[len(right.whole):]
		left.whole = left.whole[:len(right.whole)]
	}
	for left.String() != "0" || len(leftover) > 0 {
		diff := left.Deduct(right)
		if !diff.negative {
			digit++
			left = diff
		} else {
			if isReminder {
				if len(result.reminder) < 29 {
					result.reminder = fmt.Sprintf("%s%s", result.reminder, strconv.Itoa(digit))
				} else {
					break
				}
			} else {
				result.whole = fmt.Sprintf("%s%s", result.whole, strconv.Itoa(digit))
			}
			if len(leftover) == 0 {
				left.whole = fmt.Sprintf("%s%s", left.whole, "0")
				isReminder = true
			} else {
				left.whole = fmt.Sprintf("%s%s", left.whole, string(leftover[0]))
				if len(leftover) > 1 {
					leftover = leftover[1:]
				} else {
					leftover = ""
				}
			}
			digit = 0
		}
	}
	if isReminder {
		result.reminder = fmt.Sprintf("%s%s", result.reminder, strconv.Itoa(digit))
	} else {
		result.whole = fmt.Sprintf("%s%s", result.whole, strconv.Itoa(digit))
	}
	for result.whole[0] == '0' && len(result.whole) > 1 {
		result.whole = result.whole[1:]
	}
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

func (n *Number) Round(precision uint) Number {
	result := n.Copy()
	if len(result.reminder) > int(precision) {
		if digit, _ := strconv.Atoi(string(result.reminder[precision])); digit >= 5 {
			if precision > 0 {
				result = n.Add(NewNumber("0." + pad("1", PADDIRECTIONLEFT, int(precision), "0")))
			} else {
				result = n.Add(NewNumber("1"))
			}
		}
		result.reminder = result.reminder[:precision]
	}
	return result
}

func (n *Number) RoundUp(precision uint) Number {
	result := n.Copy()
	if len(result.reminder) > int(precision) {
		if precision > 0 {
			result = n.Add(NewNumber("0." + pad("1", PADDIRECTIONLEFT, int(precision), "0")))
		} else {
			result = n.Add(NewNumber("1"))
		}
		result.reminder = result.reminder[:precision]
	}
	return result
}

func (n *Number) RoundDown(precision uint) Number {
	result := n.Copy()
	if len(result.reminder) > int(precision) {
		result.reminder = result.reminder[:precision]
	}
	return result
}

func (n *Number) Ceil() Number {
	result := n.Copy()
	if len(result.reminder) > 0 {
		if digit, _ := strconv.Atoi(string(result.reminder[0])); digit >= 0 {
			result = n.Add(NewNumber("1"))
		}
		result.reminder = ""
	}
	return result
}
func (n *Number) Floor() Number {
	result := n.Copy()
	result.reminder = ""
	return result
}
