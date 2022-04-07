package usecase

import (
	"calculate/src/domain/model"
	"calculate/tools"
	"errors"
)

type factorialUsecase struct{}

type FactorialUsecase interface {
	ValidateFactorialData(*model.Factorial) error
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

func (fu *factorialUsecase) ValidateFactorialData(f *model.Factorial) error {
	if f.A > 20 || f.A < 0 || f.B > 20 || f.B < 0 {
		return errors.New("incorrect input")
	}
	return nil
}
