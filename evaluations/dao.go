package evaluations

import (
	"context"
)

//go:generate mockgen -destination=./mocks/mock_dao.go -package=mocks -source=dao.go Dao

type Dao interface {
	Save(ctx context.Context, evaluation Data) bool
	Get(ctx context.Context, tag string) bool
}


type dao struct {
	db string
}

func NewDao(db string) Dao {
	return &dao{db}
}

func (d *dao) Save(ctx context.Context, evaluation Data) bool {
	return true
}

func (d *dao) Get(ctx context.Context, tag string) bool{
	return true
}
