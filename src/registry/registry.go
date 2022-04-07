package registry

import "calculate/src/interface/controller"

type registry struct{}

type Registry interface {
	NewAPIController() controller.APIController
}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewAPIController() controller.APIController {
	return controller.APIController{
		Factorial: r.NewFactorialController(),
	}
}
