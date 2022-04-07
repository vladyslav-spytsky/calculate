package controller

type APIController struct {
	Factorial interface{ FactorialController }
}
