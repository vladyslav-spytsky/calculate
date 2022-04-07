package usecase

import (
	"calculate/src/domain/model"
	"calculate/tools"
	"errors"
)

type factorialUsecase struct{}

type FactorialUsecase interface {
	ValidateFactorialData(map[string]int) error
	CalculateFactorial(*model.Factorial) (*model.Factorial, error)
}

func NewFactorialUsecase() FactorialUsecase {
	return &factorialUsecase{}
}

func (fu *factorialUsecase) CalculateFactorial(f *model.Factorial) (*model.Factorial, error) {

	chA := make(chan int)
	chB := make(chan int)

	go tools.FactorialCalculate(f.A, chA)

	go tools.FactorialCalculate(f.B, chB)

	factorialResponse := &model.Factorial{A: <-chA, B: <-chB}

	return factorialResponse, nil
}

func (fu *factorialUsecase) ValidateFactorialData(f map[string]int) error {
	if _, ok := f["a"]; !ok {
		return errors.New("incorrect input")
	}

	if _, ok := f["b"]; !ok {
		return errors.New("incorrect input")
	}

	if len(f) > 2 {
		return errors.New("incorrect input")
	}

	if f["a"] > 20 ||f["a"] < 0 || f["b"] > 20 || f["b"] < 0 {
		return errors.New("incorrect input")
	}
	return nil
}
