package simulations

import (
	"context"
)

//go:generate mockgen -destination=./mocks/mock_dao.go -package=simulationsmocks -source=dao.go Dao

type Dao interface {
	SaveSimulation(ctx context.Context, simulation SimulationSetup)
	GetSimulation(ctx context.Context, simulationID int64)
}

type dao struct {
	db string
}

func NewDao(db string) Dao {
	return &dao{db}
}

func (s *dao) SaveSimulation(ctx context.Context, simulationSetup SimulationSetup) {

}

func (s *dao) GetSimulation(ctx context.Context, simulationID int64) {

}
