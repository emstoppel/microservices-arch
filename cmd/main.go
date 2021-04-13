package cmd

import (
	"github.com/emstoppel/microservices-arch/evaluation"
	"github.com/emstoppel/microservices-arch/simulation"
)

func main()  {
	db := "mydb"
	simGtw := simulation.NewGateway(db)
	evaluationsHandler := evaluation.NewEvaluationsService(db, simGtw)

	//route with routes.go
	evaluationsHandler.Post()

	//use my gtw
	g := evaluationsHandler.GetGateway()
	print(g)
}