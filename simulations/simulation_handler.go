package simulations

import (
	"net/http"
	"strconv"

	engineGTW "github.com/mercadolibre/fury_rules-simulator-poc/cmd/api/engine/gateway"

	"github.com/mercadolibre/go-meli-toolkit/goutils/logger"

	"github.com/gin-gonic/gin"
	rulesDomain "github.com/mercadolibre/fury_rules-simulator-poc/cmd/api/rules/domain"
	ruleGTW "github.com/mercadolibre/fury_rules-simulator-poc/cmd/api/rules/gateway"
	domain "github.com/mercadolibre/fury_rules-simulator-poc/cmd/api/simulations/domain"
	simulationGTW "github.com/mercadolibre/fury_rules-simulator-poc/cmd/api/simulations/gateway"
)

type Handler interface {
	Setup(c *gin.Context)
	GetSimulation(c *gin.Context)
	GetGateway() simulationGTW.SimulationGateway
	AddRules(c *gin.Context)
	GetRules(c *gin.Context)
	GetResults(c *gin.Context)
	Run(c *gin.Context)
}

const URL = "/simulations"

type simulationHandlerImpl struct {
	gtw       simulationGTW.SimulationGateway
	ruleGTW   ruleGTW.RuleGateway
	engineGTW engineGTW.EngineGateway
}

func NewService(simulation simulationGTW.SimulationGateway, rule ruleGTW.RuleGateway, engine engineGTW.EngineGateway) Handler {
	return &simulationHandlerImpl{
		gtw:       simulation,
		ruleGTW:   rule,
		engineGTW: engine,
	}
}

func (s *simulationHandlerImpl) GetGateway() simulationGTW.SimulationGateway {
	return s.gtw
}

func (s *simulationHandlerImpl) Setup(c *gin.Context) {
	var simulationSetup domain.SimulationSetup
	if err := c.BindJSON(&simulationSetup); err != nil {
		logger.Errorf("uri: %s Bad request unmarshalling request body", err, c.Request.RequestURI)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if err := simulationSetup.IsValid(); err != nil {
		logger.Errorf("uri: %s invalid setup", err, c.Request.RequestURI)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	simulation, err := s.gtw.SaveSimulation(c, simulationSetup)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	insertedRules, err := s.ruleGTW.InsertRules(c, simulation.ID, simulation.Setup.Rules)
	if err != nil {
		logger.Error("error inserting simulation rulesDomain", err)
		c.JSON(err.Status(), err)
		return
	}
	simulation.Setup.Rules = insertedRules
	c.JSON(http.StatusOK, simulation)
}

func (s *simulationHandlerImpl) GetSimulation(c *gin.Context) {
	simulationIDString := c.Param("id")
	simulationID, err := strconv.ParseInt(simulationIDString, 10, 64)
	if err != nil {
		logger.Errorf("uri: %s sent id is not an integer", err, c.Request.RequestURI)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	simulation, errGetSim := s.gtw.GetSimulation(c, simulationID)
	if errGetSim != nil {
		logger.Error("error", errGetSim)
		c.JSON(errGetSim.Status(), errGetSim)
		return
	}
	simulationRules, errGetRules := s.ruleGTW.GetRulesBySimulation(c, simulation.ID)
	if errGetRules != nil {
		logger.Error("error inserting simulation rulesDomain", errGetRules)
		c.JSON(errGetRules.Status(), errGetRules)
		return
	}
	simulation.Setup.Rules = simulationRules

	c.JSON(http.StatusOK, simulation)
}

func (s *simulationHandlerImpl) AddRules(c *gin.Context) {
	simulationIDString := c.Param("id")
	simulationID, err := strconv.ParseInt(simulationIDString, 10, 64)
	if err != nil {
		logger.Errorf("uri: %s sent id is not an integer", err, c.Request.RequestURI)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	var rules []rulesDomain.RuleABM
	if err := c.BindJSON(&rules); err != nil {
		logger.Errorf("uri: %s Bad request unmarshalling request body", err, c.Request.URL.Path)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	insertedRules, insertErr := s.ruleGTW.InsertRules(c, simulationID, rules)
	if insertErr != nil {
		c.JSON(insertErr.Status(), err)
		return
	}
	c.JSON(http.StatusOK, insertedRules)
}

func (s *simulationHandlerImpl) GetRules(c *gin.Context) {
	simulationIDString := c.Param("id")
	simulationID, err := strconv.ParseInt(simulationIDString, 10, 64)
	if err != nil {
		logger.Errorf("uri: %s sent id is not an integer", err, c.Request.RequestURI)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	rules, errGetRules := s.ruleGTW.GetRulesBySimulation(c, simulationID)
	if errGetRules != nil {
		logger.Error("Error getting Rules by simulation", err)
		c.JSON(errGetRules.Status(), err)
		return
	}
	c.JSON(http.StatusOK, rules)
}

func (s *simulationHandlerImpl) Run(c *gin.Context) {
	simulationIDString := c.Param("id")
	simulationID, err := strconv.ParseInt(simulationIDString, 10, 64)
	if err != nil {
		logger.Errorf("uri: %s sent id is not an integer", err, c.Request.RequestURI)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	activeSimulationID, runErr := s.engineGTW.Run(c, simulationID)
	if runErr != nil {
		logger.Errorf(runErr.Message(), runErr, c.Request.URL.Path)
		c.JSON(runErr.Status(), runErr)
		return
	}
	c.JSON(http.StatusOK, activeSimulationID)
}

func (s *simulationHandlerImpl) GetResults(c *gin.Context) {
	simulationIDString := c.Param("id")
	simulationID, err := strconv.ParseInt(simulationIDString, 10, 64)
	if err != nil {
		logger.Errorf("uri: %s sent id is not an integer", err, c.Request.RequestURI)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	executionReport, error := s.gtw.GetExecutionReport(c, simulationID)
	if error != nil {
		logger.Errorf(error.Message(), error, c.Request.URL.Path)
		c.JSON(error.Status(), error)
		return
	}
	c.JSON(http.StatusOK, executionReport)
}
