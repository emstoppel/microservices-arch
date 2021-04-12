package evaluations

import "github.com/emstoppel/microservices-arch/simulations"

//Handler
type Handler interface {
	//PostEvaluation does
	PostEvaluation()
	GetGateway() Gateway
}

const URL = "/evaluations"

type handlerImpl struct {
	gtw Gateway
}

func NewEvaluationsService(db string, simGTW simulations.Gateway) Handler {
	return &handlerImpl{
		gtw: newGateway(db, simGTW),
	}
}

func (h *handlerImpl) GetGateway() Gateway{
	return h.gtw
}

func (h *handlerImpl) PostEvaluation() {

}
