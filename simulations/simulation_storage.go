package simulations

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/mercadolibre/fury_rules-simulator-poc/cmd/api/internal/storage/mysql"
	domain "github.com/mercadolibre/fury_rules-simulator-poc/cmd/api/simulations/domain"
	"github.com/mercadolibre/go-meli-toolkit/goutils/apierrors"
)

//go:generate mockgen -destination=../../mocks/mock_simulation_storage.go -package=mocks -source=simulation_storage.go SimulationStorage

type SimulationStorage interface {
	SaveSimulation(ctx context.Context, simulation domain.SimulationSetup) (domain.Simulation, apierrors.ApiError)
	GetSimulation(ctx context.Context, simulationID int64) (domain.Simulation, apierrors.ApiError)
}

type simulationStorage struct {
	db mysql.Client
}

func NewStorage(db mysql.Client) SimulationStorage {
	return &simulationStorage{db}
}

var (
	setupQuery         = "INSERT INTO simulation (evaluation_data_type, evaluation_params) VALUE (?,?);"
	getSimulationQuery = "SELECT * FROM simulation WHERE id=?"
)

func (r *simulationStorage) SaveSimulation(ctx context.Context, simulationSetup domain.SimulationSetup) (domain.Simulation, apierrors.ApiError) {
	insertedID := int64(0)
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	err := r.db.WithTransaction(func(trx *sql.Tx) apierrors.ApiError {
		defer cancel()
		ep, err := json.Marshal(simulationSetup.EvaluationParams)
		if err != nil {
			return apierrors.NewInternalServerApiError(err.Error(), err)
		}
		inserted, err := r.db.RawExec(ctx, trx, setupQuery, simulationSetup.EvaluationDataName, string(ep))
		if err != nil {
			return apierrors.NewInternalServerApiError(err.Error(), err)
		}
		insertedID, err = inserted.LastInsertId()
		if err != nil {
			return apierrors.NewInternalServerApiError(err.Error(), err)
		}
		return nil
	})

	return domain.Simulation{ID: insertedID, Setup: simulationSetup}, err
}

func (r *simulationStorage) GetSimulation(ctx context.Context, simulationID int64) (domain.Simulation, apierrors.ApiError) {
	var evParams string
	var simulation domain.Simulation
	row := r.db.RawQueryRow(ctx, nil, getSimulationQuery, simulationID)
	err := row.Scan(&simulation.ID, &simulation.Setup.EvaluationDataName, &evParams)
	if err != nil {
		return simulation, selectError("Error getting simulation", err)
	}
	err = json.Unmarshal([]byte(evParams), &simulation.Setup.EvaluationParams)
	if err != nil {
		return domain.Simulation{}, selectError("Error getting simulation params", err)
	}
	return simulation, nil
}
func selectError(errString string, err error) apierrors.ApiError {
	if err == sql.ErrNoRows {
		return apierrors.NewNotFoundApiError(errString)
	}
	return apierrors.NewInternalServerApiError(errString, err)
}
