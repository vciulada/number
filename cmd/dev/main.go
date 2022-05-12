package main

import (
	"log"

	"github.com/vciulada/number"
)

func main() {
	log.Println("start")
	for i := 0; i < 1000000; i++ {
		a := number.NewNumber("10")
		b := number.NewNumber("12")
		_ = a.Add(b)
	}
	log.Println("end")
}
