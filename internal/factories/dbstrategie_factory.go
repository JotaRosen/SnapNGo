package factories

import (
	"snap-n-go/internal/dbstrategies"
	"snap-n-go/internal/types"
)

// Recibe parámetros y retorna la estrategia adecuada
var StrategyFactory = map[string]func(map[string]types.ConnectionParams) dbstrategies.DBStrategy{
	"mongo": func(args map[string]types.ConnectionParams) dbstrategies.DBStrategy {
		return &dbstrategies.MongoConnectionParams{
			Command:  args["cp"].Command,
			Engine:   args["cp"].Engine,
			Host:     args["cp"].Host,
			Port:     args["cp"].Port,
			Username: args["cp"].Username,
			Password: args["cp"].Password,
			DbName:   args["cp"].DbName,
		}
	},
}
