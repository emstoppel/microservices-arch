package evaluation

import (
	evaluationsmocks "github.com/emstoppel/microservices-arch/evaluation/mocks"
	simulationsmocks "github.com/emstoppel/microservices-arch/simulation/mocks"
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
	//init mocked simulation gtw
	mockedSimulationGtw := simulationsmocks.NewMockGateway(mockCtrl)

	// Initialize mocked gtw
	evaluationsGateway := &gateway{
		simGTW: mockedSimulationGtw,
		dao:    mockedDao,
	}

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	evaluationsGateway.Get(c, "TAG")
}