package evaluations

type Handler interface {
	PostEvaluation()
}

const URL = "/evaluations"

type handlerImpl struct {
	gtw Gateway
}

func NewEvaluationService(db string) Handler{
	return &handlerImpl{
		gtw:NewGateway(db),
	}
}

func (h *handlerImpl) PostEvaluation() {

}
