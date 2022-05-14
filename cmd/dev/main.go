package main

import (
	"fmt"
	"log"

	"github.com/vciulada/number"
)

func main() {
	log.Println("start")
	a := number.NewNumber("144")
	fmt.Println(a.Devide(number.NewNumber("12")))
	log.Println("end")
}
