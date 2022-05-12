package main

import (
	"fmt"
	"log"

	"github.com/vciulada/number"
)

func main() {
	log.Println("start")
	for i := 0; i < 100000; i++ {
		a := number.NewNumber("10")
		b := number.NewNumber("12")
		_ = a.Add(b)
		fmt.Println(a.Add(b).String())
	}
	log.Println("end")
}
