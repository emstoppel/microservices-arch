package cmd

import (
	"github.com/emstoppel/microservices-arch/evaluations"
	"github.com/emstoppel/microservices-arch/simulations"
)

func main()  {
	db := "mydb"
	simGtw := simulations.NewGateway(db)
	evaluationsHandler := evaluations.NewEvaluationsService(db, simGtw)

	//route with routes.go
	evaluationsHandler.PostEvaluation()

	//use my gtw
	evaluationsHandler.GetGateway()
}