package evaluations

import (
	"context"
	"github.com/emstoppel/microservices-arch/simulations"
)

//go:generate mockgen -destination=./mocks/mock_gateway.go -package=evaluationsmocks -source=gateway.go Gateway

// Always comment wat Gateway is
type Gateway interface {
	// Always comment what CreateEvaluation does
	CreateEvaluation(ctx context.Context, evaluation EvaluationData)
	//Always comment what GetEvaluations does, you can refer to other functions
	//like CreateEvaluation
	GetEvaluations(ctx context.Context, tag string)
}

type gateway struct {
	simGTW simulations.Gateway
	dao    Dao
}

func NewGateway(db string, simGTW simulations.Gateway) Gateway {
	return &gateway{dao: NewDao(db), simGTW: simGTW}
}

func (r *gateway) CreateEvaluation(ctx context.Context, evaluation EvaluationData) {
	//maybe do smth with simGTW ;)
}

func (r *gateway) GetEvaluations(ctx context.Context, tag string) {
	//maybe do smth with simGTW ;)
}
