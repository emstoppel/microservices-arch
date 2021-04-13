package simulation

import (
	"context"
)

//go:generate mockgen -destination=./mocks/mock_dao.go -package=simulationsmocks -source=dao.go Dao

type Dao interface {
	Save(ctx context.Context, simulation SimulationSetup)
	Get(ctx context.Context, simulationID int64)
}

type dao struct {
	db string
}

func NewDao(db string) Dao {
	return &dao{db}
}

func (s *dao) Save(ctx context.Context, simulationSetup SimulationSetup) {

}

func (s *dao) Get(ctx context.Context, simulationID int64) {

}
