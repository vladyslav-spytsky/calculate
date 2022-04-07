package controller

import (
	"calculate/src/domain/model"
	"calculate/src/usecase"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type factorialController struct {
	factorialUsecase usecase.FactorialUsecase
}

type FactorialController interface {
	ValidateFactorialData(*http.Request) error
	CalculateFactorial(http.ResponseWriter, *http.Request, httprouter.Params)
}

func NewFactorialController(fu usecase.FactorialUsecase) FactorialController {
	return &factorialController{fu}
}

func (fc *factorialController) CalculateFactorial(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	f := &model.Factorial{}
	_ = json.NewDecoder(r.Body).Decode(f)

	resp, err := fc.factorialUsecase.CalculateFactorial(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respJson, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(respJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func (fc *factorialController) ValidateFactorialData(r *http.Request) error {

	f := &model.Factorial{}

	err := json.NewDecoder(r.Body).Decode(f)
	if err != nil {
		return errors.New("incorrect input")

	}

	err = fc.factorialUsecase.ValidateFactorialData(f)
	if err != nil {
		return errors.New("incorrect input")
	}

	return nil
}
