package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/vciulada/number"
)

func main() {
	log.Println("start")
	for i := 0; i < 1000000; i++ {
		a := number.NewNumber(strconv.Itoa(i))
		number := a.Multiply(a)
		fmt.Println(number)
	}
	log.Println("end")
}
