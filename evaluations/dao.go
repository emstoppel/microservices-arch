package evaluations

import (
	"context"
)

//go:generate mockgen -destination=./mocks/mock_dao.go -package=evaluationsmocks -source=dao.go Dao

//dao lives cause
type Dao interface {
	//Save does
	Save(ctx context.Context, evaluation EvaluationData) bool
	//Get does
	Get(ctx context.Context, tag string) bool
}

type dao struct {
	db string
}

func NewDao(db string) Dao {
	return &dao{db}
}

func (d *dao) Save(ctx context.Context, evaluation EvaluationData) bool {
	return true
}

func (d *dao) Get(ctx context.Context, tag string) bool {
	return true
}
