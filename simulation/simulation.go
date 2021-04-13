package simulation

import (
	"errors"

	"github.com/emstoppel/microservices-arch/site"
)

var (
	ErrInvalidSite = errors.New("the simulation has an invalid site")
)

type Simulation struct {
	ID     int64            `json:"id"`
	Setup  SimulationSetup  `json:"setup"`
	Result SimulationResult `json:"result"`
}

type SimulationSetup struct {
	//always payments
	Entity string `json:"entity"`
	Site   string `json:"site"`
	//always processing fee id
	RuleSet int64 `json:"rule_set"`
	//evaluationData (DS) name
	EvaluationDataName string `json:"evaluation_data_type"`
	//execution params (POC nil)
	EvaluationParams `json:"evaluation_params"`
}
type SimulationResult map[string]interface{}

type EvaluationParams map[string]interface{}
type SimulationRules map[string]interface{}

func (s SimulationSetup) IsValid() error {
	if !site.Exists(s.Site) {
		return ErrInvalidSite
	}
	return nil
}
