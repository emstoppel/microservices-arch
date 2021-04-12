package evaluations

import (
	evaluationsmocks "github.com/emstoppel/microservices-arch/evaluations/mocks"
	simulationsmocks "github.com/emstoppel/microservices-arch/simulations/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"testing"
)

func TestMockSimulationGateways(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//init mocked dao
	mockedDao := evaluationsmocks.NewMockDao(mockCtrl)
	//init mocked simulations gtw
	mockedSimulationGtw := simulationsmocks.NewMockGateway(mockCtrl)

	// Initialize mocked gtw
	evaluationsGateway := &gateway{
		simGTW: mockedSimulationGtw,
		dao:    mockedDao,
	}

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	evaluationsGateway.GetEvaluations(c, "TAG")
}