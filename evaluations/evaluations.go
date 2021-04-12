package evaluations

type EvaluationInput map[string]interface{}
type RuleSetOutput map[string]interface{}

type EvaluationData struct {
	ID              string `json:"id"`
	EvaluationInput `json:"evaluation_input"`
	EvaluationResult
	Tags []string `json:"tags"`
}

type EvaluationResult struct {
	Results []Result `json:"evaluation_result"`
}

type Result struct {
	RuleSetName  string        `json:"rule_set_name"`
	RuleTypeName string        `json:"rule_type_name"`
	RuleID       int64         `json:"rule_id"`
	Output       RuleSetOutput `json:"output"`
}
