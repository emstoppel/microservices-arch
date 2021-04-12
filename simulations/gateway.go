package simulations

import (
	"context"
)

//go:generate mockgen -destination=./mocks/mock_gateway.go -package=simulationsmocks -source=gateway.go Gateway

//A
type Gateway interface {
	//SAVE
	SaveSimulation(ctx context.Context, request SimulationSetup)
	//GET
	GetSimulation(ctx context.Context, simulationID int64)
}

type gateway struct {
	storage Dao
}

func NewGateway(db string) Gateway {
	return &gateway{storage: NewDao(db)}
}

// Setup will set the params to execute the simulation.
func (simulation *gateway) SaveSimulation(ctx context.Context, request SimulationSetup) {

}

func (simulation *gateway) GetSimulation(ctx context.Context, simulationID int64) {
}
