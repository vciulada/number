package number

import (
	"fmt"
	"log"
	"strconv"
	"strings"
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

func (n *Number) Add(a Number) Number {
	result := Number{}
	j := len(a.whole) - 1

	var memory int
	for i := len(n.whole) - 1; i >= 0; i-- {
		x, _ := strconv.Atoi(string(n.whole[i]))
		y, _ := strconv.Atoi(string(a.whole[j]))
		sum := strconv.Itoa(x + y + memory)
		memory = 0
		var digit string
		if len(sum) > 1 {
			memory, _ = strconv.Atoi(string(sum[0]))
			digit = string(sum[1])
		} else {
			digit = sum
		}
		result.whole = fmt.Sprintf("%s%s", digit, result.whole)
		j--
		if j < 0 && i > 0 {
			for z := i - 1; z >= 0; z-- {
				x, _ := strconv.Atoi(string(n.whole[z]))
				sum := strconv.Itoa(x + memory)
				memory = 0
				var digit string
				if len(sum) > 1 {
					memory, _ = strconv.Atoi(string(sum[0]))
					digit = string(sum[1])
				} else {
					digit = sum
				}
				result.whole = fmt.Sprintf("%s%s", digit, result.whole)
			}
			break
		}
	}
	if j >= 0 {
		for z := j; z >= 0; z-- {
			x, _ := strconv.Atoi(string(a.whole[z]))
			sum := strconv.Itoa(x + memory)
			memory = 0
			var digit string
			if len(sum) > 1 {
				memory, _ = strconv.Atoi(string(sum[0]))
				digit = string(sum[1])
			} else {
				digit = sum
			}
			result.whole = fmt.Sprintf("%s%s", digit, result.whole)
		}
	}
	if memory > 0 {
		digit := strconv.Itoa(memory)
		result.whole = fmt.Sprintf("%s%s", digit, result.whole)
	}

	return result
}
