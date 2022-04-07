package registry

import (
	interfaceController "calculate/src/interface/controller"
	usecaseUsecase "calculate/src/usecase"
)

func (r *registry) NewFactorialController() interfaceController.FactorialController {
	return interfaceController.NewFactorialController(r.NewFactorialUsecase())
}

func (r *registry) NewFactorialUsecase() usecaseUsecase.FactorialUsecase {
	return usecaseUsecase.NewFactorialUsecase()
}
