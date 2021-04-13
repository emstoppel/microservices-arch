package evaluation

import (
	"context"
	"github.com/emstoppel/microservices-arch/simulation"
)

//go:generate mockgen -destination=./mocks/mock_gateway.go -package=evaluationsmocks -source=gateway.go Gateway

// Always comment wat Gateway is
type Gateway interface {
	// Always comment what CreateEvaluation does
	Create(ctx context.Context, evaluation EvaluationData)
	//Always comment what GetEvaluations does, you can refer to other functions
	//like CreateEvaluation
	Get(ctx context.Context, tag string)
}

type gateway struct {
	simGTW simulation.Gateway
	dao    Dao
}

func newGateway(db string, simGTW simulation.Gateway) Gateway {
	return &gateway{dao: NewDao(db), simGTW: simGTW}
}

func (r *gateway) Create(ctx context.Context, evaluation EvaluationData) {
	//maybe do smth with simGTW ;)
}

func (r *gateway) Get(ctx context.Context, tag string) {
	//maybe do smth with simGTW ;)
}
