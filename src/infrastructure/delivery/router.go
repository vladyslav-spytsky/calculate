package delivery

import (
	"calculate/src/interface/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(router *httprouter.Router, c controller.APIController) *httprouter.Router {

	router.POST("/calculate", validateFactorialMiddleware(c.Factorial.CalculateFactorial, c))

	return router
}
