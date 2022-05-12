package number

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	PADDIRECTIONRIGHT = iota
	PADDIRECTIONLEFT
)

func NewNumber(number string) Number {
	result := Number{}
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
	whole    string
	reminder string
}

func (n Number) String() string {
	if n.reminder == "" {
		return n.whole
	}
	return fmt.Sprintf("%s.%s", n.whole, n.reminder)
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
func addReminder(pLeft, pRight string) (string, int) {
	left := pad(pLeft, PADDIRECTIONRIGHT, len(pRight), "0")
	right := pad(pRight, PADDIRECTIONRIGHT, len(pLeft), "0")

	result := ""

	var memory int
	for i := len(left) - 1; i >= 0; i-- {
		x, _ := strconv.Atoi(string(left[i]))
		y, _ := strconv.Atoi(string(right[i]))
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

func addWhole(pLeft, pRight string, pMemory int) string {
	left := pad(pLeft, PADDIRECTIONLEFT, len(pRight), "0")
	right := pad(pRight, PADDIRECTIONLEFT, len(pLeft), "0")

	result := ""

	memory := pMemory
	for i := len(left) - 1; i >= 0; i-- {
		x, _ := strconv.Atoi(string(left[i]))
		y, _ := strconv.Atoi(string(right[i]))
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
	if memory > 0 {
		digit := strconv.Itoa(memory)
		result = fmt.Sprintf("%s%s", digit, result)
	}

	return result
}

func (n *Number) Add(a Number) Number {

	result := Number{}
	var memory int
	result.reminder, memory = addReminder(n.reminder, a.reminder)
	result.whole = addWhole(n.whole, a.whole, memory)

	return result
}
