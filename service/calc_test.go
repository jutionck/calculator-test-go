package service

import (
	"log"
	"testing"
)

func TestCalculator_Add(t *testing.T) {
	t.Run("Calculator Add operator testing", func(t *testing.T) {
		numSample1 := 10
		numSample2 := 20
		calc := Calculator{
			Num1: float64(numSample1),
			NUm2: float64(numSample2),
		}
		r, err := calc.Add()
		if err != nil {
			log.Fatalln(err)
		}
		if *r != 30 {
			t.Error("It should return 30")
		}
	})

	t.Run("Calculator detected negative number", func(t *testing.T) {
		numSample1 := -10
		numSample2 := 20
		calc := Calculator{
			Num1: float64(numSample1),
			NUm2: float64(numSample2),
		}
		_, err := calc.Add()
		if err != nil && err.Error() != "negative number detected" {
			t.Error("It have a negative number")
		}
	})
}

func TestCalculator_Sub(t *testing.T) {
	t.Run("Calculator Sub operator testing", func(t *testing.T) {
		numSample1 := 3
		numSample2 := 2

		calc := Calculator{
			Num1: float64(numSample1),
			NUm2: float64(numSample2),
		}

		r, err := calc.Sub()
		if err != nil {
			log.Fatalln(err)
		}
		if *r != 1 {
			t.Error("It should return 1")
		}
	})

	t.Run("Calculator detected negative number", func(t *testing.T) {
		numSample1 := -10
		numSample2 := 20
		calc := Calculator{
			Num1: float64(numSample1),
			NUm2: float64(numSample2),
		}
		_, err := calc.Sub()
		if err != nil && err.Error() != "negative number detected" {
			t.Error("It have a negative number")
		}
	})
}

/*
Run :
go test calculator-test/service

go test -v calculator-test/service -coverprofile=cover.out && go tool cover -html=cover.out
*/
