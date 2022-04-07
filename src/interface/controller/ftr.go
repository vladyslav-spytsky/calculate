package controller

import (
	"calculate/src/domain/model"
	"calculate/src/usecase"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type factorialController struct {
	factorialUsecase usecase.FactorialUsecase
}

type FactorialController interface {
	ValidateFactorialData(http.ResponseWriter,*http.Request) error
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

func (fc *factorialController) ValidateFactorialData(w http.ResponseWriter, r *http.Request) error {

	var g map[string]int

	if errors := json.NewDecoder(r.Body).Decode(&g); errors != nil {
		err := map[string]interface{}{"error": "incorrect input"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return errors
	}

	if errors := fc.factorialUsecase.ValidateFactorialData(g); errors != nil {
		err := map[string]interface{}{"error": errors.Error()}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return errors
	}

	return nil
}
