package main

import (
	"calculator-test/service"
	"flag"
	"fmt"
)

func main() {
	num1 := flag.Int("num1", 0, "first number")
	num2 := flag.Int("num2", 0, "first number")
	opr := flag.String("opr", "add", "Calculator operator")
	flag.Parse()

	switch *opr {
	case "add":
		myCalc := service.Calculator{
			Num1: *num1, NUm2: *num2,
		}
		fmt.Println(myCalc.Add())
	case "sub":
		myCalc := service.Calculator{
			Num1: *num1, NUm2: *num2,
		}
		fmt.Println(myCalc.Sub())
	}

}

/*
Running :
go run main/app.go -num1 1 -num2 3 -opr add
*/
