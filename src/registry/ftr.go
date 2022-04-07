package registry

import (
	interfaceController "calculate/src/interface/controller"
	Usecase "calculate/src/usecase"
)

func (r *registry) NewFactorialController() interfaceController.FactorialController {
	return interfaceController.NewFactorialController(r.NewFactorialUsecase())
}

func (r *registry) NewFactorialUsecase() Usecase.FactorialUsecase {
	return Usecase.NewFactorialUsecase()
}
