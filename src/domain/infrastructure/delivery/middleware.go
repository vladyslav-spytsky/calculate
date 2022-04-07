package delivery

import (
	"calculate/src/interface/controller"
	"calculate/tools"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Validate data factorial.
func validateFactorialMiddleware(handle httprouter.Handle, c controller.APIController) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		r2 := tools.CopyHttpRequest(r)
		err := c.Factorial.ValidateFactorialData(r2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handle(w, r, p)
	}
}
