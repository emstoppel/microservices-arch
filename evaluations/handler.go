package evaluations

import "github.com/emstoppel/microservices-arch/simulations"

//Handler
type Handler interface {
	//PostEvaluation does
	PostEvaluation()
}

const URL = "/evaluations"

type handlerImpl struct {
	gtw Gateway
}

func NewEvaluationService(db string, simGTW simulations.Gateway) Handler {
	return &handlerImpl{
		gtw: NewGateway(db, simGTW),
	}
}

func (h *handlerImpl) PostEvaluation() {

}
