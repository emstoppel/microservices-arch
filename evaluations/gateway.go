package evaluations

import (
	"context"
)

//go:generate mockgen -destination=./mocks/mock_gtw.go -package=mocks -source=gateway.go Gateway

type Gateway interface {
	CreateEvaluation(ctx context.Context, evaluation Data)
	GetEvaluations(ctx context.Context, tag string)
}

type service struct {
	storage Dao
}

func NewGateway(db string) Gateway {
	return &service{storage: NewDao(db)}
}

func (r *service) CreateEvaluation(ctx context.Context, evaluation Data) {
}

func (r *service) GetEvaluations(ctx context.Context, tag string) {
}
