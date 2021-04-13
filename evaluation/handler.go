package evaluation

import "github.com/emstoppel/microservices-arch/simulation"

//Handler
type Handler interface {
	//Post does
	Post()
	GetGateway() Gateway
}

const URL = "/evaluation"

type handlerImpl struct {
	gtw Gateway
}

func NewEvaluationsService(db string, simGTW simulation.Gateway) Handler {
	return &handlerImpl{
		gtw: newGateway(db, simGTW),
	}
}

func (h *handlerImpl) GetGateway() Gateway{
	return h.gtw
}

func (h *handlerImpl) Post() {

}
