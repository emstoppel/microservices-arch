package simulations

import (
	"context"

	domainResults "github.com/mercadolibre/fury_rules-simulator-poc/cmd/api/results/domain"
	resultsGTW "github.com/mercadolibre/fury_rules-simulator-poc/cmd/api/results/gateway"
	"github.com/mercadolibre/go-meli-toolkit/goutils/apierrors"

	"github.com/mercadolibre/fury_rules-simulator-poc/cmd/api/internal/storage/mysql"
	domain "github.com/mercadolibre/fury_rules-simulator-poc/cmd/api/simulations/domain"
	storage "github.com/mercadolibre/fury_rules-simulator-poc/cmd/api/simulations/storage"
)

//go:generate mockgen -destination=../../mocks/mock_simulation_gtw.go -package=mocks -source=simulation_gtw.go SimulationGateway

type SimulationGateway interface {
	SaveSimulation(ctx context.Context, request domain.SimulationSetup) (domain.Simulation, apierrors.ApiError)
	GetSimulation(ctx context.Context, simulationID int64) (domain.Simulation, apierrors.ApiError)
	GetExecutionReport(ctx context.Context, simulationID int64) (domainResults.ExecutionReport, apierrors.ApiError)
}

type simulationService struct {
	storage   storage.SimulationStorage
	resultGTW resultsGTW.ResultGateway
}

func NewServiceGateway(db mysql.Client, gateway resultsGTW.ResultGateway) SimulationGateway {
	return &simulationService{storage: storage.NewStorage(db), resultGTW: gateway}
}

// Setup will set the params to execute the simulation.
func (simulation *simulationService) SaveSimulation(ctx context.Context, request domain.SimulationSetup) (domain.Simulation, apierrors.ApiError) {
	return simulation.storage.SaveSimulation(ctx, request)
}

func (simulation *simulationService) GetSimulation(ctx context.Context, simulationID int64) (domain.Simulation, apierrors.ApiError) {
	return simulation.storage.GetSimulation(ctx, simulationID)
}

func (simulation *simulationService) GetExecutionReport(ctx context.Context, simulationID int64) (domainResults.ExecutionReport, apierrors.ApiError) {
	return simulation.resultGTW.GetExecutionReport(ctx, simulationID)
}
