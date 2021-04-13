package simulation

import (
	"context"
)

//go:generate mockgen -destination=./mocks/mock_gateway.go -package=simulationsmocks -source=gateway.go Gateway

//A
type Gateway interface {
	//SAVE
	Save(ctx context.Context, request SimulationSetup)
	//GET
	Get(ctx context.Context, simulationID int64)
}

type gateway struct {
	storage Dao
}

func NewGateway(db string) Gateway {
	return &gateway{storage: NewDao(db)}
}

// Setup will set the params to execute the simulation.
func (simulation *gateway) Save(ctx context.Context, request SimulationSetup) {

}

func (simulation *gateway) Get(ctx context.Context, simulationID int64) {
}
